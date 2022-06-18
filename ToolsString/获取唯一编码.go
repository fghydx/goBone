package GLStrings

import (
	"io"
	"crypto/rand"
	"encoding/base64"
	"MyLib/GLCrypto"
)

func GetGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GLCrypto.Md5Hex(base64.URLEncoding.EncodeToString(b))
}