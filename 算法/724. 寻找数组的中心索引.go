package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 2, 3}))
}

/*
输入：
nums = [1, 7, 3, 6, 5, 6]
输出：3
解释：
索引 3 (nums[3] = 6) 的左侧数之和 (1 + 7 + 3 = 11)，与右侧数之和 (5 + 6 = 11) 相等。
同时, 3 也是第一个符合要求的中心索引。
*/

func pivotIndex(nums []int) int {
	leftSum := 0
	sum := 0
	for _, num := range nums {
		sum += num
	}
	for i := 0; i < len(nums); i++ {

		if leftSum == sum - nums[i] - leftSum{
			return i
		} else {
			leftSum += nums[i]
		}

	}
	return -1
}
