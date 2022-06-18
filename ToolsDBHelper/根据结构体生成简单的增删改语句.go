package ToolsDBHelper

import (
	"fmt"
	"reflect"
	"strings"
)

//简单的生成SQL语句
type DbStruct struct {
}

//type Testt struct {
//	DbStruct
//	Aid int `col:"aid"`
//	Abc string `col:"abc"`
//	Bcd int `col:"bcd"`
//	Cde float64 `col:"cde"`
//}
//生成插入语句
func (DS *DbStruct) Insert(TblName string, obj interface{}) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if len(TblName) == 0 {
		TblName = t.Name()
	}
	fieldNum := t.NumField()
	tmpstr := "insert into %s(%s) values(%s)"
	field := ""
	values := ""
	for i := 0; i < fieldNum; i++ {
		ftag := t.Field(i).Tag.Get("col")
		if len(ftag) == 0 {
			continue
		}
		tmp := strings.Split(ftag, ",")
		omitempty := false
		if len(tmp) == 2 {
			omitempty = tmp[1] == "omitempty"
		}
		field = field + tmp[0] + ","

		switch v.Field(i).Kind() {
		case reflect.String:
			if !(omitempty && v.Field(i).String() == "") {
				values = values + fmt.Sprintf(`"%s",`, v.Field(i).String())
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if !(omitempty && v.Field(i).Int() == 0) {
				values = values + fmt.Sprintf(`%d,`, v.Field(i).Int())
			}
		case reflect.Float32, reflect.Float64:
			if !(omitempty && v.Field(i).Float() == 0) {
				values = values + fmt.Sprintf(`%f,`, v.Field(i).Float())
			}
		case reflect.Bool:
			if !(omitempty && !v.Field(i).Bool()) {
				values = values + fmt.Sprintf(`%t,`, v.Field(i).Bool())
			}
		default:
			if !(omitempty && len(v.Field(i).Bytes()) == 0) {
				values = values + string(v.Field(i).Bytes()) + ","
			}
		}
	}
	return fmt.Sprintf(tmpstr, TblName, field[:len(field)-1], values[:len(values)-1])
}

func (DS *DbStruct) GenerateArgs(obj interface{}) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	field := ""
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		ftag := t.Field(i).Tag.Get("col")
		if len(ftag) == 0 {
			continue
		}
		tmp := strings.Split(ftag, ",")
		omitempty := false
		if len(tmp) == 2 {
			omitempty = tmp[1] == "omitempty"
		}
		colName := "@" + tmp[0]
		switch v.Field(i).Kind() {
		case reflect.String:
			if !(omitempty && v.Field(i).String() == "") {
				field = field + colName + fmt.Sprintf(`="%s",`, v.Field(i).String())
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if !(omitempty && v.Field(i).Int() == 0) {
				field = field + colName + fmt.Sprintf(`=%d,`, v.Field(i).Int())
			}
		case reflect.Float32, reflect.Float64:
			if !(omitempty && v.Field(i).Float() == 0) {
				field = field + colName + fmt.Sprintf(`=%f,`, v.Field(i).Float())
			}
		case reflect.Bool:
			if !(omitempty && !v.Field(i).Bool()) {
				field = field + colName + fmt.Sprintf(`=%t,`, v.Field(i).Bool())
			}
		default:
			if !(omitempty && len(v.Field(i).Bytes()) == 0) {
				field = field + colName + "=" + string(v.Field(i).Bytes()) + ","
			}
		}
	}
	return field[:len(field)-1]
}

func (DS *DbStruct) Update(TblName string, obj interface{}, Key string) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	field := ""
	where := ""
	tmpstr := "update %s set %s where %s"
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		ftag := t.Field(i).Tag.Get("col")
		if len(ftag) == 0 {
			continue
		}
		tmp := strings.Split(ftag, ",")
		omitempty := false
		if len(ftag) == 2 {
			omitempty = tmp[1] == "omitempty"
		}
		colName := tmp[0]
		switch v.Field(i).Kind() {
		case reflect.String:
			if colName == Key {
				where = Key + fmt.Sprintf(`="%s"`, v.Field(i).String())
				continue
			}
			if !(omitempty && v.Field(i).String() == "") {
				field = field + colName + fmt.Sprintf(`="%s",`, v.Field(i).String())
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if colName == Key {
				where = Key + fmt.Sprintf(`=%d`, v.Field(i).Int())
				continue
			}
			if !(omitempty && v.Field(i).Int() == 0) {
				field = field + colName + fmt.Sprintf(`=%d,`, v.Field(i).Int())
			}
		case reflect.Float32, reflect.Float64:
			if colName == Key {
				where = Key + fmt.Sprintf(`=%f`, v.Field(i).Float())
				continue
			}
			if !(omitempty && v.Field(i).Float() == 0) {
				field = field + colName + fmt.Sprintf(`=%f,`, v.Field(i).Float())
			}
		case reflect.Bool:
			if colName == Key {
				where = Key + fmt.Sprintf(`=%t`, v.Field(i).Bool())
				continue
			}
			if !(omitempty && !v.Field(i).Bool()) {
				field = field + colName + fmt.Sprintf(`=%t,`, v.Field(i).Bool())
			}
		default:
			if colName == Key {
				where = Key + "=" + string(v.Field(i).Bytes())
				continue
			}
			if !(omitempty && len(v.Field(i).Bytes()) == 0) {
				field = field + colName + "=" + string(v.Field(i).Bytes()) + ","
			}
		}
	}
	return fmt.Sprintf(tmpstr, TblName, field[:len(field)-1], where)
}

func (DS *DbStruct) Delete(TblName string, obj interface{}, Key string) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	where := ""
	tmpstr := "delete %s where %s"
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		if len(where) > 0 {
			break
		}
		ftag := t.Field(i).Tag.Get("col")
		if len(ftag) == 0 {
			continue
		}
		tmp := strings.Split(ftag, ",")
		colName := tmp[0]
		if colName == Key {
			switch v.Field(i).Kind() {
			case reflect.String:
				where = Key + fmt.Sprintf(`="%s"`, v.Field(i).String())
				break
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				where = Key + fmt.Sprintf(`=%d`, v.Field(i).Int())
				break
			case reflect.Float32, reflect.Float64:
				where = Key + fmt.Sprintf(`=%f`, v.Field(i).Float())
				break
			case reflect.Bool:
				where = Key + fmt.Sprintf(`=%t`, v.Field(i).Bool())
				break
			default:
				where = Key + "=" + string(v.Field(i).Bytes())
				break
			}
		}
	}
	return fmt.Sprintf(tmpstr, TblName, where)
}
