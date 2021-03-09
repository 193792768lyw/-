package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findShortestSubArray([]int{1}))
}

func findShortestSubArray(nums []int) int {
	type Index struct {
		Nums int
		Min  int
		Max  int
	}
	info := make(map[int]*Index, 0)
	du := 1
	for key, value := range nums {
		if va, ok := info[value]; ok {
			va.Nums += 1
			if va.Nums > du {
				du = va.Nums
			}
			va.Max = key
		} else {
			info[value] = &Index{
				Nums: 1,
				Min:  key,
				Max:  key,
			}
		}
	}

	res := math.MaxInt32
	for _, value := range info {
		if value.Nums == du {
			if value.Max-value.Min+1 < res {
				res = value.Max - value.Min + 1
			}
		}
	}
	return res
}
