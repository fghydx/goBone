package ToolsString

import (
	"regexp"
	"strings"
)

// DeleteAllSpace 利用正则表达式压缩字符串，去除空格或制表符
func DeleteAllSpace(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

// StrToSlice 将字符串按指定字符串打散成字符串切片
func StrToSlice(str string, split string) (result []string) {
	if str == "" {
		return nil
	}
	if split == "" {
		return append(result, str)
	}
	reg := regexp.MustCompile(split + "+")
	tmp := reg.ReplaceAllString(str, split)
	result = strings.Split(tmp, split)
	return
}

// ReverseString 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
