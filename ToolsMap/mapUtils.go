package ToolsMap

// KeysAndValues 返回map的所有Key和Value
func KeysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	var values []V
	for _, k := range keys {
		values = append(values, m[k])
	}
	return keys, values
}
