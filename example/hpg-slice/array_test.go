package main

import (
	"fmt"
	"testing"
)

func TestArrayAssign(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := [4]int{2, 4, 5, 6}
	fmt.Println(a, b)
	// a = b // cannot use b (type [4]int) as type [3]int in assignment
}

/*
Go 语言中，数组变量属于值类型(value type)，因此当一个数组变量被赋值或者传递时，实际上会复制整个数组。
例如，将 a 赋值给 b，修改 a 中的元素并不会改变 b 中的元素：
*/
func TestArrayCopy(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := a
	a[0] = 100
	fmt.Println(a, b)
}

func square(arr *[3]int) {
	for i, num := range *arr {
		(*arr)[i] = num * num
	}
}

func TestArrayPointer(t *testing.T) {
	a := [...]int{1, 2, 3}
	square(&a)
	fmt.Println(a)
	if a[1] != 4 && a[2] != 9 {
		t.Fatal("failed")
	}
}
