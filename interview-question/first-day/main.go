package main

import "fmt"

func main1() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

/*
参考解析：这是新手常会犯的错误写法，for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，
所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3.
*/

func main() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		value := val
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "===>", *v)
	}
}
