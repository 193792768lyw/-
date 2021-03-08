package main

import (
	"fmt"
)

func main() {
	fmt.Println(minKBitFlips([]int{0,0,0,1,0,1,1,0},3))
}
func minKBitFlips(A []int, K int) int {
	res := 0
	for i := 0 ;i <= len(A) - K ; i++{
		if A[i] == 0 {
			for j := i ; j < len(A) && j < i + K ; j++{
				A[j] = A[j] ^ 1
			}
			res++
		}
	}
	for _, value := range A{
		if value == 0  {
			return -1
		}
	}
	return res
}