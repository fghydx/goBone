package ToolsSyncObj

import (
	"MyLib/GLFile"
	"MyLib/ThirdUnit"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

// Synchronous FIFO queue
type SyncQueue struct {
	lock    sync.Mutex
	popable *sync.Cond
	buffer  *ThirdUnit.Queue
	closed  bool
}

// Create a new SyncQueue
func NewSyncQueue() *SyncQueue {
	ch := &SyncQueue{
		buffer: ThirdUnit.NewQueue(),
	}
	ch.popable = sync.NewCond(&ch.lock)
	return ch
}

// Pop an item from SyncQueue, will block if SyncQueue is empty
func (q *SyncQueue) Pop() (v interface{}) {
	c := q.popable
	buffer := q.buffer

	q.lock.Lock()
	defer q.lock.Unlock()
	for buffer.Length() == 0 && !q.closed {
		c.Wait()
	}

	if buffer.Length() > 0 {
		v = buffer.Peek()
		buffer.Remove()
	}
	return
}

//
func (q *SyncQueue) TryPop() (v interface{}) {
	buffer := q.buffer

	q.lock.Lock()
	defer q.lock.Unlock()
	if buffer.Length() > 0 {
		v = buffer.Peek()
		buffer.Remove()
	} else if q.closed {
		v = nil
	}
	return
}

// Push an item to SyncQueue. Always returns immediately without blocking
func (q *SyncQueue) Push(v interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if !q.closed {
		q.buffer.Add(v)
		q.popable.Signal()
	}
}

// Get the length of SyncQueue
func (q *SyncQueue) Len() (l int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	l = q.buffer.Length()
	return
}

// Close SyncQueue
//
// After close, Pop will return nil without block, and TryPop will return v=nil, ok=True
func (q *SyncQueue) Close() {
	q.lock.Lock()
	defer q.lock.Unlock()
	if !q.closed {
		q.closed = true
		q.popable.Broadcast()
	}
}

func (q *SyncQueue) Closed() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.closed {
		return true
	}
	return false
}

func (q *SyncQueue) SaveToFile(filename string) {
	GLFile.ClearFile(filename)
	buffer := q.buffer
	q.lock.Lock()
	defer q.lock.Unlock()
	for {
		if buffer.Length() > 0 {
			v := buffer.Peek()
			buffer.Remove()
			str, err := json.Marshal(&v)
			if err != nil {
				fmt.Println(err.Error() + "\r\n")
				return
			}
			GLFile.AppendTextToFile(filename, string(str)+"\r\n")
		} else {
			return
		}
	}
}

func (q *SyncQueue) LoadFromFile(filename string) {
	filestr := GLFile.ReadTextFromFile(filename)
	str := strings.TrimSpace(string(filestr))
	datas := strings.Split(str, "\r\n")
	q.lock.Lock()
	defer q.lock.Unlock()
	for _, value := range datas {
		var data interface{}
		err := json.Unmarshal([]byte(value), &data)
		if err != nil {
			fmt.Println("SyncQueue.LoadFromFile", err.Error()+"\r\n")
			continue
		}
		q.Push(data)
	}
}
