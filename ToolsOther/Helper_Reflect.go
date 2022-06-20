package ToolsOther

import "reflect"

// CreateObjFromObj 根据已知的类型,创建一个对像,其实相当于创建了一个相同类型的对像,参数其实是个对像
//Type XXXX struct{}
//aa := CreateObjFromType(XXXX{})
//反回值  aa.(*XXXX) 转换
func CreateObjFromObj(ImplObj interface{}) interface{} {
	t := reflect.ValueOf(ImplObj).Type()
	v := reflect.New(t).Interface()
	return v
}

// GetObjType 获到对像的类型
func GetObjType(obj interface{}) reflect.Type {
	t := reflect.ValueOf(obj).Type()
	return t
}

// GetObjRefType 获到对像的类型地址
func GetObjRefType(obj interface{}) reflect.Type {
	t := reflect.Indirect(reflect.ValueOf(obj)).Type()
	return t
}

func CreateObjFromType(typ reflect.Type) interface{} {
	v := reflect.New(typ).Interface()
	return v
}

