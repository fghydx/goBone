package ToolsCache

import (
	"encoding/json"
	"fmt"
	ToolsFiles "gobone/ToolsFiles"
	"gobone/ToolsOther"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type SessionData struct {
	Key      interface{}
	Token    string
	LastTime int64
	RefCount int64 //状态
	Age      int64 //单位为秒
	Data     interface{}
}

type TOnDelete = func(interface{})
type TOnLoadOneFromFile = func(interface{}, interface{}) //Session,Key

type SessionManage struct {
	MaxAge      int64 //单位为秒
	locker      sync.RWMutex
	OnTimeOuted TOnDelete
	OnLoadOne   TOnLoadOneFromFile
	Data        map[interface{}]SessionData
}

func (AData *SessionManage) Range(f func(key, value interface{}) bool) {
	AData.locker.Lock()
	defer AData.locker.Unlock()
	for key, value := range AData.Data {
		if !f(key, value.Data) {
			break
		}
	}
}

func (AData *SessionManage) SaveToFile(filename string) {
	ToolsFiles.ClearFile(filename)
	AData.locker.Lock()
	defer AData.locker.Unlock()
	for _, value := range AData.Data {
		str, err := json.Marshal(&value.Data)
		if err != nil {
			fmt.Println("K-V超时管理单元:SaveToFile" + err.Error() + "\r\n")
			return
		}
		ToolsFiles.AppendTextToFile(filename, string(str)+"\r\n")
		value.Data = nil
		str, err = json.Marshal(&value)
		if err != nil {
			fmt.Println("K-V超时管理单元:SaveToFile" + err.Error() + "\r\n")
			return
		}
		ToolsFiles.AppendTextToFile(filename, string(str)+"\r\n")
	}
}

func (AData *SessionManage) LoadFromFile(filename string, Adata interface{}) {
	if b, _ := ToolsFiles.Exists(filename); !b {
		return
	}
	str := string(ToolsFiles.ReadTextFromFile(filename))
	str = strings.TrimSpace(str)
	if str == "" {
		return
	}
	datas := strings.Split(str, "\r\n")
	AData.locker.Lock()
	defer AData.locker.Unlock()
	var data = ToolsOther.CreateObjFromObj(Adata)
	for index, value := range datas {
		if (index+1)%2 != 0 {
			err := json.Unmarshal([]byte(value), data)
			if err != nil {
				fmt.Println("K-V超时管理单元:LoadFromFile" + err.Error() + "\r\n")
				continue
			}
		} else {
			var sessionData = new(SessionData)
			err := json.Unmarshal([]byte(value), sessionData)
			if err != nil {
				fmt.Println("K-V超时管理单元:LoadFromFile" + err.Error() + "\r\n")
				continue
			}
			sessionData.LastTime = time.Now().Unix()
			sessionData.Data = data
			AData.Data[sessionData.Key] = *sessionData
			if AData.OnLoadOne != nil {
				AData.OnLoadOne(data, sessionData.Key)
			}
			if index+1 < len(datas) {
				data = ToolsOther.CreateObjFromObj(Adata)
			}

		}
	}
}

func (AData *SessionManage) timeOutCheck() {
	AData.locker.Lock()
	defer func() {
		time.AfterFunc(time.Second*3, func() {
			AData.timeOutCheck()
		})
		AData.locker.Unlock()
	}()

	for key, value := range AData.Data {
		CurTime := time.Now().Unix()
		if atomic.LoadInt64(&value.RefCount) > 0 {
			continue
		}

		if value.Age > 0 { //如果有单个超时设置(Age>0),用单个超时,否则判断全局超时设置
			if value.LastTime+value.Age <= CurTime {
				if AData.OnTimeOuted != nil {
					AData.OnTimeOuted(value.Data)
				}
				delete(AData.Data, key)
			}
		} else if AData.MaxAge > 0 { //如果有全局超时设置(MaxAge>0),用全局超时,否则永不超时
			if value.LastTime+AData.MaxAge <= CurTime {
				if AData.OnTimeOuted != nil {
					AData.OnTimeOuted(value.Data)
				}
				delete(AData.Data, key)
			}
		}
	}
}

func (AData *SessionManage) Count() int {
	AData.locker.Lock()
	defer AData.locker.Unlock()

	return len(AData.Data)
}

func NewSessionManage(timeOut int64, AOnTimeOuted TOnDelete, AOnLoadOneFromFile TOnLoadOneFromFile) *SessionManage {
	AData := &SessionManage{MaxAge: timeOut, OnTimeOuted: AOnTimeOuted, OnLoadOne: AOnLoadOneFromFile, Data: make(map[interface{}]SessionData)}
	go AData.timeOutCheck()
	return AData
}

func (AData *SessionManage) AddSessionData(Key interface{}, Token string, Data interface{}, Age int64) {
	tmpData := SessionData{Key, Token, time.Now().Unix(), 0, Age, Data}
	AData.locker.Lock()
	defer AData.locker.Unlock()
	AData.Data[Key] = tmpData
}

func (AData *SessionManage) CheckSession(Key interface{}, Token string, Del bool, UpdateLastTime bool) (Exist bool, valid bool) {
	AData.locker.RLock()
	defer AData.locker.RUnlock()
	Res, Finded := AData.Data[Key]
	if Finded {
		if Res.Token == Token {
			if Del {
				if AData.OnTimeOuted != nil {
					AData.OnTimeOuted(Res.Data)
				}
				delete(AData.Data, Key)
			} else {
				if UpdateLastTime {
					Res.LastTime = time.Now().Unix()
				}
			}
			return true, true
		} else {
			return true, false
		}
	}
	return false, false
}

func (AData *SessionManage) DelSession(Key interface{}, GiveBack bool) {
	AData.locker.Lock()
	defer AData.locker.Unlock()
	Res, Finded := AData.Data[Key]
	if Finded {
		if GiveBack {
			atomic.AddInt64(&Res.RefCount, -1)
		}
		if atomic.LoadInt64(&Res.RefCount) <= 0 {
			if AData.OnTimeOuted != nil {
				AData.OnTimeOuted(Res.Data)
			}
			delete(AData.Data, Key)
		}

	}
}

//NeedGiveBack为true时需要调用GiveBack
func (AData *SessionManage) GetSessionData(Key interface{}, UpdateLastTime bool, NeedGiveBack bool) interface{} {
	if Key == "" {
		return nil
	}
	AData.locker.RLock()
	defer AData.locker.RUnlock()
	Res, Finded := AData.Data[Key]
	if Finded {
		if UpdateLastTime {
			Res.LastTime = time.Now().Unix()
		}
		if NeedGiveBack {
			atomic.AddInt64(&Res.RefCount, 1)
		}
		return Res.Data
	}
	return nil
}

func (AData *SessionManage) GiveBack(Key interface{}) {
	AData.locker.RLock()
	defer AData.locker.RUnlock()
	Res, Finded := AData.Data[Key]
	if Finded {
		atomic.AddInt64(&Res.RefCount, -1)
	}
}

func (AData *SessionManage) UpdateSession(Key interface{}) {
	AData.locker.Lock()
	defer AData.locker.Unlock()
	Res, Finded := AData.Data[Key]
	if Finded {
		Res.LastTime = time.Now().Unix()
	}
}

func (AData *SessionManage) Clear() {
	AData.locker.Lock()
	defer AData.locker.Unlock()
	AData.Data = make(map[interface{}]SessionData)
}
