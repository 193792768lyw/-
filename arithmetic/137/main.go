package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
}

/*
输入：nums = [0,1,0,1,0,1,99]
输出：99
*/
func singleNumber1(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			i += 3
		} else {
			return nums[i]
		}
	}
	return 0
}

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
