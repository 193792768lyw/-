package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(11, 4))
}
func hammingDistance(x int, y int) int {
	res := 0
	for i := 0; i < 31; i++ {
		res += (x >> i & 0x0001) ^ (y >> i & 0x0001)
	}
	return res
}

func hammingDistance1(x, y int) (ans int) {
	for s := x ^ y; s > 0; s >>= 1 {
		ans += s & 1
	}
	return
}
