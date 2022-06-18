package ToolsSlice

import "testing"

func TestSliceInsert(t *testing.T) {
	slice := []string{"1", "2", "3", "4", "5", "6"}
	_, slice = SliceInsert(slice, "7a", 3)
	t.Logf("%v", slice)
}

func TestSliceInsertEx(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice1 := []int{1, 2, 3, 4, 5, 6}
	_, slice = SliceInsertEx(slice, slice1, 3)
	t.Logf("%v", slice)
}

func TestSliceDelete(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	_, slice = SliceDelete(slice, 3)
	t.Logf("%v", slice)
}

func TestSliceMoveToEnd(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	_, slice = SliceMoveToEnd(slice, 3)
	t.Logf("%v", slice)
}

func TestSliceMoveToBegin(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	_, slice = SliceMoveToBegin(slice, 3)
	t.Logf("%v", slice)
}

func TestRemoveRepByLoop(t *testing.T) {
	slice := []string{"1", "2", "3", "4", "5", "6", "3", "2"}
	slice = RemoveRepByLoop(slice)
	t.Logf("%v", slice)
}

func TestRemoveRepByMap(t *testing.T) {
	slice := []string{"1", "2", "3", "4", "5", "6", "3", "2"}
	slice = RemoveRepByMap(slice)
	t.Logf("%v", slice)
}

func TestContainsValue(t *testing.T) {
	slice := []string{"1", "2", "3", "4", "5", "6", "3", "2"}
	b := ContainsValue(slice, "1")
	t.Logf("%v", b)
}
