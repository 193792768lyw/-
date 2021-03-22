package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(0b00000000000000000000000000001011))
}
func hammingWeight1(num uint32) int {
	sum := 0
	for num != 0 {
		if num&1 != 0 {
			sum++
		}
		num >>= 1
	}
	return sum
}
func hammingWeight(num uint32) (ones int) {
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			ones++
		}
	}
	return
}
