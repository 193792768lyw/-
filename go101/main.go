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
	a := [3]int{-1, 0, 1}
	s := []bool{true, false}
	m := map[string]int{"abc": 123, "xyz": 789}
	fmt.Println(a[2], s[1], m["abc"])     // 读取
	a[2], s[1], m["abc"] = 999, true, 567 // 修改
	fmt.Println(a[2], s[1], m["abc"])     // 读取

	n, present := m["hello"]
	fmt.Println(n, present, m["hello"]) // 0 false 0
	n, present = m["abc"]
	fmt.Println(n, present, m["abc"]) // 567 true 567
	m = nil
	fmt.Println(m["abc"]) // 0

	// 下面这两行编译不同过。
	/*
		_ = a[3]  // 下标越界
		_ = s[-1] // 下标越界
	*/

	// 下面这几行每行都会造成一个恐慌。
	_ = a[n]         // panic: 下标越界
	_ = s[n]         // panic: 下标越界
	m["hello"] = 555 // panic: m为一个零值映射
}
