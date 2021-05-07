package main

import "fmt"

func main() {
	fmt.Println(xorOperation(1, 7))
}
func xorOperation(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res ^= start + i*2
	}
	return res
}
