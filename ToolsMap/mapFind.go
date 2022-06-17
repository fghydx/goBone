package ToolsMap

// KeysAndValues 返回map的所有Key和Value
func KeysAndValues(m map[interface{}]interface{}) ([]interface{}, []interface{}) {
	var keys []interface{}
	for k := range m {
		keys = append(keys, k)
	}
	var values []interface{}
	for _, k := range keys {
		values = append(values, m[k])
	}
	return keys, values
}
