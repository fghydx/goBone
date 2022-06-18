package GLStrings

import "math/rand"
import "time"

func init() {
	rand.Seed(time.Now().UnixNano())
}

type StringType int

const (
	Letter StringType = iota
	Num
	All
)

var letterRunes0 = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var letterRunes1 = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var letterRunes2 = []rune("1234567890")

func RandStringRunes(n int,t StringType) string {
	b := make([]rune, n)
	switch t {
	case Letter:
		for i := range b {
			b[i] = letterRunes1[rand.Intn(len(letterRunes1))]
		}
	case Num:
		for i := range b {
			b[i] = letterRunes2[rand.Intn(len(letterRunes2))]
		}
	case All:
		for i := range b {
			b[i] = letterRunes0[rand.Intn(len(letterRunes0))]
		}
	}

	return string(b)
}

func RandInt(n int) int {
	return rand.Intn(n)
}
