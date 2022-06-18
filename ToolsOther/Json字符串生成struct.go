package ToolsOther

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func createSubStruct(m map[string]interface{}, structName string) (result []string) {
	var buffer bytes.Buffer
	buffer.WriteString("type ")
	buffer.WriteString(structName)
	buffer.WriteString(" struct {\n")
	for k, v := range m {
		runes := []rune(k)
		buffer.WriteString("	")
		fname := strings.ToUpper(string(runes[0])) + string(runes[1:])
		buffer.WriteString(fname)
		buffer.WriteString("   ")
		//buffer.WriteString(fmt.Sprintf("%T", v))
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			if value, ok := v.([]interface{})[0].(map[string]interface{}); ok {
				substruct := createSubStruct(value, "T"+fname)
				result = append(result, substruct...)
				buffer.WriteString("[]T" + fname)
			}
			if value, ok := v.([]interface{}); ok {
				buffer.WriteString("[]" + reflect.TypeOf(value[0]).String())
			}
			buffer.WriteString("     `json:\"")
			buffer.WriteString(k)
			buffer.WriteString("\"`\n")
			continue
		case reflect.Map:
			if value, ok := v.(map[string]interface{}); ok {
				substruct := createSubStruct(value, "T"+fname)
				result = append(result, substruct...)
				buffer.WriteString("T" + fname)
			}
		default:
			buffer.WriteString(reflect.TypeOf(v).String())
		}
		buffer.WriteString("     `json:\"")
		buffer.WriteString(k)
		buffer.WriteString("\"`")
		buffer.WriteString("	// ")
		buffer.WriteString(fmt.Sprintf("%v", reflect.ValueOf(v).Interface()))
		buffer.WriteString("\n")
	}
	buffer.WriteString("}")
	result = append(result, buffer.String())
	return
}

func JsonToStruct(jsonStr string, structName string) (result []string) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println("转化错误:", err)
		return nil
	}
	result = createSubStruct(m, structName)
	return
}
