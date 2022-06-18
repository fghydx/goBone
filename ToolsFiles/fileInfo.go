// golang 实现读取exe dll 版本号
package ToolFiles

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	MZ       = "MZ"
	PE       = "PE"
	RSRC     = ".rsrc"
	TYPET    = 16
	PEOFFSET = 64
	MACHINE  = 332
	DEFAULT  = `C:\Windows\System32\cmd.exe`
)

type fileInfo struct {
	FilePath string
	versionA []int64
	version  string
}

func GetFileVersion(filepath string) string {
	fileinfo := &fileInfo{FilePath: filepath}
	fileinfo.GetExeVersion()
	return fileinfo.version
}

func GetFileVersionA(filepath string) []int64 {
	fileinfo := &fileInfo{FilePath: filepath}
	fileinfo.GetExeVersion()
	return fileinfo.versionA
}

func GetDirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func GetFileSize(path string) int64 {
	if b, _ := Exists(path); !b {
		return 0
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

func GetFileModTime(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		return time.Unix(0, 0).Unix()
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Unix(0, 0).Unix()
	}
	return fi.ModTime().Unix()
}

func (f *fileInfo) checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// 获取exe dll版本
func (f *fileInfo) GetExeVersion() (err error) {
	file, err := os.Open(f.FilePath)
	f.checkError(err)

	// 第一次读取64 byte
	buffer := make([]byte, 64)
	_, err = file.Read(buffer)
	f.checkError(err)
	defer file.Close()

	str := string(buffer[0]) + string(buffer[1])
	if str != MZ {
		log.Fatalln("读取exe错误,找不到 MZ.", f.FilePath)
	}

	peOffset := f.unpack([]byte{buffer[60], buffer[61], buffer[62], buffer[63]})
	if peOffset < PEOFFSET {
		log.Fatalln("peOffset 读取错误.", f.FilePath)
	}

	// 读取从文件开头移位到 peOffset，第二次读取 24 byte
	_, err = file.Seek(int64(peOffset), 0)
	buffer = make([]byte, 24)
	_, err = file.Read(buffer)
	f.checkError(err)

	str = string(buffer[0]) + string(buffer[1])
	if str != PE {
		log.Fatalln("读取exe错误,找不到 PE.", f.FilePath)
	}

	machine := f.unpack([]byte{buffer[4], buffer[5]})
	if machine != MACHINE {
		log.Fatalln("machine 读取错误.", f.FilePath)
	}

	noSections := f.unpack([]byte{buffer[6], buffer[7]})
	optHdrSize := f.unpack([]byte{buffer[20], buffer[21]})

	// 读取从当前位置移位到 optHdrSize，第三次读取 40 byte
	file.Seek(int64(optHdrSize), 1)
	resFound := false
	for i := 0; i < int(noSections); i++ {
		buffer = make([]byte, 40)
		file.Read(buffer)
		str = string(buffer[0]) + string(buffer[1]) + string(buffer[2]) + string(buffer[3]) + string(buffer[4])
		if str == RSRC {
			resFound = true
			break
		}
	}
	if !resFound {
		log.Fatalln("读取exe错误,找不到 .rsrc.", f.FilePath)
	}

	infoVirt := f.unpack([]byte{buffer[12], buffer[13], buffer[14], buffer[15]})
	infoSize := f.unpack([]byte{buffer[16], buffer[17], buffer[18], buffer[19]})
	infoOff := f.unpack([]byte{buffer[20], buffer[21], buffer[22], buffer[23]})

	// 读取从文件开头位置移位到 infoOff，第四次读取 infoSize byte
	file.Seek(int64(infoOff), 0)
	buffer = make([]byte, infoSize)
	_, err = file.Read(buffer)
	f.checkError(err)

	nameEntries := f.unpack([]byte{buffer[12], buffer[13]})
	idEntries := f.unpack([]byte{buffer[14], buffer[15]})

	var infoFound bool
	var subOff, i int64
	for i = 0; i < (nameEntries + idEntries); i++ {
		typeT := f.unpack([]byte{buffer[i*8+16], buffer[i*8+17], buffer[i*8+18], buffer[i*8+19]})
		if typeT == TYPET {
			infoFound = true
			subOff = int64(f.unpack([]byte{buffer[i*8+20], buffer[i*8+21], buffer[i*8+22], buffer[i*8+23]}))
			break
		}
	}
	if !infoFound {
		log.Fatalln("读取exe错误,找不到 typeT == 16.", f.FilePath)
	}

	subOff = subOff & 0x7fffffff
	infoOff = f.unpack([]byte{buffer[subOff+20], buffer[subOff+21], buffer[subOff+22], buffer[subOff+23]}) //offset of first FILEINFO
	infoOff = infoOff & 0x7fffffff
	infoOff = f.unpack([]byte{buffer[infoOff+20], buffer[infoOff+21], buffer[infoOff+22], buffer[infoOff+23]}) //offset to data
	dataOff := f.unpack([]byte{buffer[infoOff], buffer[infoOff+1], buffer[infoOff+2], buffer[infoOff+3]})
	dataOff = dataOff - infoVirt

	version1 := f.unpack([]byte{buffer[dataOff+48], buffer[dataOff+48+1]})
	version2 := f.unpack([]byte{buffer[dataOff+48+2], buffer[dataOff+48+3]})
	version3 := f.unpack([]byte{buffer[dataOff+48+4], buffer[dataOff+48+5]})
	version4 := f.unpack([]byte{buffer[dataOff+48+6], buffer[dataOff+48+7]})

	version := fmt.Sprintf("%d.%d.%d.%d", version2, version1, version4, version3)
	f.versionA = append(f.versionA, version2)
	f.versionA = append(f.versionA, version1)
	f.versionA = append(f.versionA, version4)
	f.versionA = append(f.versionA, version3)
	f.version = version

	return nil
}

func (f *fileInfo) unpack(b []byte) (num int64) {
	for i := 0; i < len(b); i++ {
		num = 256*num + int64((b[len(b)-1-i] & 0xff))
	}
	return
}
