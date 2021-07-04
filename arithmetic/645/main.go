package main

import (
	"fmt"
)

func main() {
	fmt.Println(findErrorNums([]int{3, 2, 3, 4, 6, 5}))
}
func findErrorNums(nums []int) []int {
	numsTemp := make([]int, len(nums))
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		numsTemp[nums[i]-1] += 1
		if numsTemp[nums[i]-1] > 1 {
			res = append(res, nums[i])
		}
	}

	for i := range numsTemp {
		if numsTemp[i] == 0 {
			res = append(res, i+1)
		}
	}
	return res
}
