package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {

	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[res] {
			continue
		} else {
			res++
			nums[res] = nums[i]
		}
	}
	return res + 1
}
