package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{11, 13, 15, 17}))
}
func findMin(nums []int) int {
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		if res > nums[i] {
			return nums[i]
		}
	}
	return res
}
