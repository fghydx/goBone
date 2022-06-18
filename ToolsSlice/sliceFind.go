package ToolsSlice

// ContainsValue 切片中是否包含元素
func ContainsValue[T comparable](slice []T, s T) bool {
	for _, s2 := range slice {
		if s2 == s {
			return true
		}
	}
	return false
}
