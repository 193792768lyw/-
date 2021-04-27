package main

import (
	"fmt"
	_ "math/rand" // okay: 匿名引入
)

type S0 struct {
	y int "foo"
	x bool
}

type S1 = struct { // S1是一个非定义类型
	x int "foo"
	y bool
}

type S2 = struct { // S2也是一个非定义类型
	x int "bar"
	y bool
}

type S3 S2 // S3是一个定义类型。
type S4 S3 // S4是一个定义类型。
// 如果不考虑字段标签，S3（S4）和S1的底层类型一样。
// 如果考虑字段标签，S3（S4）和S1的底层类型不一样。

var v0, v1, v2, v3, v4 = S0{}, S1{}, S2{}, S3{}, S4{}

func main() {
	arr := []int{1, 1, 1, 1, 1, 1, 1}
	arr1 := arr[0:2:3]
	arr1 = append(arr1, 9, 9)
	fmt.Println(arr1, len(arr1), cap(arr1))
	fmt.Println(arr)
}
