package ToolFiles

import (
	"io/ioutil"
	"os"
)

///追加写
func AppendTextToFile(fileName, text string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		println(err.Error())
		return err
	}
	defer file.Close()

	//	file.Seek(0,os.SEEK_END)
	file.WriteString(text)
	return nil
}

///追加写
func AppendDataToFile(fileName string, data []byte) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		println(err.Error())
		return err
	}
	defer file.Close()

	//	file.Seek(0,os.SEEK_END)
	file.Write(data)
	return nil
}

///覆盖写
func WriteTextToFile(fileName, text string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(text)
	return nil
}

func ReadTextFromFile(FileName string) []byte {
	fi, err := os.Open(FileName)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}
	return fd
}
