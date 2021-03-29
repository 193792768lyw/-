package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(reverseBits(0b11111111111111111111111111111101))
}
func reverseBits(num uint32) uint32 {
	var by strings.Builder
	for i := 0; i < 32; i++ {
		if num&1 == 0 {
			by.WriteString("0")
		} else {
			by.WriteString("1")
		}
		num = num >> 1
	}
	//fmt.Println(by.String())
	res, _ := strconv.ParseUint(by.String(), 2, 32)
	return uint32(res)

}
