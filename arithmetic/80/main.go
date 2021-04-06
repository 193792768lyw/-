package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
}

/*
输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
*/
func removeDuplicates(nums []int) int {
	index := 0
	count := 1
	for _, num := range nums {
		if index != 0 && nums[index-1] != num {
			count = 1
		}
		if count <= 2 {
			nums[index] = num
			count++
			index++
		}

	}
	return index
}
