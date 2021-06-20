package main

import (
	"reflect"
	"testing"
)

func TestSplitArray(t *testing.T) {
	res := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9, 10}}
	if !reflect.DeepEqual(SplitArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3), res) {
		t.Log("错误")
	}
}
