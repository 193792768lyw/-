package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}

// 154. 寻找旋转排序数组中的最小值 II
func findMin(nums []int) int {
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		if res > nums[i] {
			return nums[i]
		}
	}
	return res
}
