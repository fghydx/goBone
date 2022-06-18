package ToolsConvert

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"time"
	"unsafe"
)

//整形32位转换成字节
func Int32ToBytes_LittleEndian(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

//整形32位转换成字节
func Int32ToBytes_BigEndian(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形32位
func BytesToInt32_LittleEndian(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return int(x)
}
func BytesToInt32_BigEndian(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func Int64ToBytes_BigEndian(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}
func Int64ToBytes_LittleEndian(i int64) []byte {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64_BigEndian(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
func BytesToInt64_LittleEndian(buf []byte) int64 {
	return int64(binary.LittleEndian.Uint64(buf))
}

func Float32ToByte_LittleEndian(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func Float32ToByte_BigEndian(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, bits)

	return bytes
}

func ByteToFloat32_LittleEndian(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}
func ByteToFloat32_BigEndian(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

func Float64ToByte_LittleEndian(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}
func Float64ToByte_BigEndian(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bits)

	return bytes
}

func ByteToFloat64_LittleEndian(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}
func ByteToFloat64_BigEndian(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

//浮点数保留几位小数
func FormatFloat(value float64, decNum uint8) (result float64) {
	result, err := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(int(decNum))+"f", value), 64)
	if err != nil {
		panic(err)
	}
	return result
}

//字符串转浮点型
func StrToFloat(value string) (result float64) {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return result
}
func StrToFloatDef(value string, defvalue float64) (result float64) {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		result = defvalue
	}
	return result
}

func StrToInt(value string) (result int) {
	result, _ = strconv.Atoi(value)
	return
}

func StrToIntdef(value string, defvalue int) (result int) {
	result, err := strconv.Atoi(value)
	if err != nil {
		result = defvalue
	}
	return
}

func StrToDateTime(str string, fmt string) time.Time {
	result, err := time.Parse(fmt, str)
	if err != nil {
		return time.Unix(0, 0)
	}
	return result
}

func UnixTimeToDateTime(value int64) time.Time {
	return time.Unix(value, 0)
}

func IntToStr(value int) (result string) {
	return strconv.Itoa(value)
}

func Int64ToStr(value int64) (result string) {
	return strconv.FormatInt(value, 10)
}

func SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToSliceByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
