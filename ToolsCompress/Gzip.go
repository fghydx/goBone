package GLCompress

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"fmt"
)

func Compress_Gzip(value []byte) []byte{
	buf := new(bytes.Buffer)
	gr := gzip.NewWriter(buf)
	defer gr.Close()
	gr.Write(value)
	gr.Flush()
	return buf.Bytes()
}

func UnCompress_Gzip(value []byte) []byte{
	buf := new(bytes.Buffer)
	buf.Write(value)
	gr,err := gzip.NewReader(buf)
	if err == nil {
		defer gr.Close()
		bufout,err := ioutil.ReadAll(gr)
		if err == nil {
			return bufout
		} else {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				return bufout
			} else {
				fmt.Println("Gzip单元:UnCompress_Gzip Error", err.Error())
				return value
			}
		}
	} else {
		fmt.Println("Gzip单元:UnCompress_Gzip Error",err.Error())
		return value
	}
}