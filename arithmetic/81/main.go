package main

import "fmt"

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
}

/*
输入：nums = [2,5,6,0,0,1,2], target = 3
输出：false
*/
func search(nums []int, target int) bool {
	len := len(nums)
	for _, num := range nums {
		if num == target {
			return true
		}
		if num > target {
			break
		}
	}

	for i := len - 1; i >= 0; i-- {
		if nums[i] == target {
			return true
		}
		if nums[i] < target {
			break
		}
	}
	return false
}
