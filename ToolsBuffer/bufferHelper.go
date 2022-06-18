package GLBuffer

import (
	"bytes"
	"encoding/binary"
)

type Buffer struct {
	bytes.Buffer
}

func (b *Buffer) ReadInt8(BigEndian bool) (Result int8) {
	if BigEndian {
		binary.Read(b, binary.BigEndian, &Result)
	} else {
		binary.Read(b, binary.LittleEndian, &Result)
	}
	return
}
func (b *Buffer) WriteInt8(v int8,BigEndian bool) {
	if BigEndian {
		binary.Write(b, binary.BigEndian, &v)
	} else {
		binary.Write(b, binary.LittleEndian, &v)
	}
}

func (b *Buffer) ReadInt16(BigEndian bool) (Result int16) {
	if BigEndian {
		binary.Read(b, binary.BigEndian, &Result)
	} else {
		binary.Read(b, binary.LittleEndian, &Result)
	}
	return
}
func (b *Buffer) WriteInt16(v int16,BigEndian bool) {
	if BigEndian {
		binary.Write(b, binary.BigEndian, &v)
	} else {
		binary.Write(b, binary.LittleEndian, &v)
	}
}

func (b *Buffer) ReadInt32(BigEndian bool) (Result int32) {
	if BigEndian {
		binary.Read(b, binary.BigEndian, &Result)
	} else {
		binary.Read(b, binary.LittleEndian, &Result)
	}
	return
}
func (b *Buffer) WriteInt32(v int32,BigEndian bool) {
	if BigEndian {
		binary.Write(b, binary.BigEndian, &v)
	} else {
		binary.Write(b, binary.LittleEndian, &v)
	}
}

func (b *Buffer) ReadInt64(BigEndian bool) (Result int64) {
	if BigEndian {
		binary.Read(b, binary.BigEndian, &Result)
	} else {
		binary.Read(b, binary.LittleEndian, &Result)
	}
	return
}
func (b *Buffer) WriteInt64(v int64,BigEndian bool) {
	if BigEndian {
		binary.Write(b, binary.BigEndian, &v)
	} else {
		binary.Write(b, binary.LittleEndian, &v)
	}
}

func (b *Buffer) ReadFloat32(BigEndian bool) (Result float32) {
	if BigEndian {
		binary.Read(b, binary.BigEndian, &Result)
	} else {
		binary.Read(b, binary.LittleEndian, &Result)
	}
	return
}
func (b *Buffer) WriteFloat32(v float32,BigEndian bool) {
	if BigEndian {
		binary.Write(b,binary.BigEndian,&v)
	} else {
		binary.Write(b,binary.LittleEndian,&v)
	}
}

func (b *Buffer) ReadFloat64(BigEndian bool) (Result float64) {
	if BigEndian {
		binary.Read(b, binary.BigEndian, &Result)
	} else {
		binary.Read(b, binary.LittleEndian, &Result)
	}
	return
}
func (b *Buffer) WriteFloat64(v float64,BigEndian bool) {
	if BigEndian {
		binary.Write(b,binary.BigEndian,&v)
	} else {
		binary.Write(b,binary.LittleEndian,&v)
	}
}

func (b *Buffer) ReadString() (Result string) {
	len := b.ReadInt32(true)
	res := make([]byte,len)
	b.Read(res)
	Result = string(res)
	return
}
func (b *Buffer) WriteString(v string) {
	len := int32(len(v))
	b.WriteInt32(len,true)
	b.Write([]byte(v))
}

func (b *Buffer) ReadRune() (Result []rune) {
	len := b.ReadInt32(true)
	res := make([]byte,len)
	b.Read(res)
	Result = bytes.Runes(res)
	return
}
func (b *Buffer) WriteRune(v []rune) {
	brune := []byte(string(v))
	len := int32(len(brune))
	b.WriteInt32(len,true)
	b.Write(brune)
}