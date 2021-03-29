package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseInt("99", 10, 64))
	if "-5" < "0" {
		fmt.Println("vdibfhn")
	}
	//fmt.Println(hammingWeight(0b00000000000000000000000000001011))
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
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
