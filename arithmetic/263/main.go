package main

import "fmt"

// 263. 丑数
func main() {
	fmt.Println(isUgly(0))
}
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	arr := []int{2, 3, 5}
	for _, v := range arr {
		for n%v == 0 {
			n /= v
		}
		if n == 1 {
			return true
		}
	}
	return false
}
