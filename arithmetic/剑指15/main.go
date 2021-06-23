package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(hammingWeight(9))
}

func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}
