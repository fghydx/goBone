package ToolsSlice

import (
	"errors"
)

// SliceInsert 向指定位置插入值
func SliceInsert[T any](slice []T, v T, index int) (err error, c []T) {
	length := len(slice)
	if index > length {
		return errors.New("切片长度<插入值index"), nil
	}
	c = append(c, slice[:index]...)
	c = append(c, v)
	c = append(c, slice[index:]...)
	return nil, c
}

// SliceInsertEx 向指定位置插入slice
func SliceInsertEx[T any](slice []T, v []T, index int) (err error, c []T) {
	length := len(slice)
	if index > length {
		return errors.New("切片长度<插入值index"), nil
	}
	c = append(c, slice[:index]...)
	c = append(c, v...)
	c = append(c, slice[index:]...)
	return nil, c
}

// SliceDelete 删除指定位置的值
func SliceDelete[T any](slice []T, index int) (err error, c []T) {
	length := len(slice)
	if index > length {
		return errors.New("切片长度<index"), nil
	}
	c = append(c, slice[:index]...)
	c = append(c, slice[index+1:]...)
	return nil, c
}

// SliceMoveToEnd 将指定位置移到最后面
func SliceMoveToEnd[T any](slice []T, index int) (err error, c []T) {
	length := len(slice)
	if index > length {
		return errors.New("切片长度<index"), nil
	}
	v := slice[index]
	c = append(c, slice[:index]...)
	c = append(c, slice[index+1:]...)
	c = append(c, v)
	return nil, c
}

// SliceMoveToBegin 将指定位置移到最前面
func SliceMoveToBegin[T any](slice []T, index int) (err error, c []T) {
	length := len(slice)
	if index > length {
		return errors.New("切片长度<index"), nil
	}
	c = append(c, slice[index])
	c = append(c, slice[:index]...)
	c = append(c, slice[index+1:]...)
	return nil, c
}

// RemoveRepByLoop 双重循环去重，切片中数据少时可用这种
func RemoveRepByLoop[T comparable](slc []T) []T {
	var result []T // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// RemoveRepByMap 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap[T comparable](stringSlice []T) []T {
	keys := make(map[T]bool)
	var list []T
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
