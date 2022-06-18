package GLFile

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GetCurrentPath 获取当前路径
func GetCurrentPath() string {
	return filepath.Dir(os.Args[0]) + string(filepath.Separator)
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreatePath(path string) bool {
	if b, _ := Exists(path); b {
		return true
	} else {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			println("GLFiles单元:CreatePath", err.Error())
			return false
		}
		return true
	}
}

func ReadFile(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return fd, nil
}

func WriteFile(path string, file []byte) error {
	return ioutil.WriteFile(path, file, 0655)
}

func GetFilelist(path string) ([]string, error) {
	var list []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		list = append(list, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}

func ClearFile(fileName string) error {
	return os.Truncate(fileName, 0)
}

func DeleteFile(fileName string) bool {
	return os.Remove(fileName) == nil
}

func DeletePath(pathName string) bool {
	return os.RemoveAll(pathName) == nil
}

func ReadBufio(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}

	bufReader := bufio.NewReader(file)
	buf := make([]byte, fi.Size())

	readedNub := int64(0)
	for {
		readNum, err := bufReader.Read(buf[readedNub:])
		if err != nil && err != io.EOF {
			panic(err)
		} else {
			if err == io.EOF {
				break
			}
		}
		readedNub = readedNub + int64(readNum)
		if fi.Size() == readedNub {
			break
		}
	}
	return buf
}

func RedirectStdOut(fileName string) {
	f, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND,
		0755)
	os.Stdout = f
}
func RedirectStderr(fileName string) {
	f, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND,
		0755)
	os.Stderr = f
}

func CopyFile(dest, source string) error {
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	f, err := os.Open(source)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(df, f)
	return err
}
