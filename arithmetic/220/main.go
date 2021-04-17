package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
}

/*
输入：nums = [1,2,3,1], k = 3, t = 0
输出：true
*/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for k1, _ := range nums {
		start := 0
		end := 0
		if k1-k <= 0 {
			start = 0
		} else {
			start = k1 - k
		}

		if k1+k >= len(nums) {
			end = len(nums) - 1
		} else {
			end = k1 + k
		}

		for j := start; j < end; j++ {
			if j != k1 && int(math.Abs(float64(nums[j]-nums[k1]))) <= t {
				return true
			}
		}
	}
	return false
}
