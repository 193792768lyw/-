package main

import "fmt"

func main() {
	fmt.Println(totalHammingDistance([]int{4, 14, 2}))
}

func totalHammingDistance(nums []int) int {
	res := 0
	for i := 0; i <= 29; i++ {
		cou0 := 0
		for _, v := range nums {
			cou0 += v >> i & 1
		}
		res += cou0 * (len(nums) - cou0)
	}

	return res
}
