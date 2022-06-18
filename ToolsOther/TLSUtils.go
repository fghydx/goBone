package ToolsOther

import (
	"encoding/pem"
	"github.com/golang/crypto/pkcs12"
	GLFile "gobone/ToolsFiles"
)

//将pfx格试的签名转换成pem,转后certFile, keyFile，都传生成后的这个pem文件
func PfxToPem(PfxFile string, Key string, SavePemFile string) {
	p12, _ := GLFile.ReadFile(PfxFile)
	blocks, err := pkcs12.ToPEM(p12, Key)
	if err != nil {
		panic(err)
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	GLFile.WriteFile(SavePemFile, pemData)
}
