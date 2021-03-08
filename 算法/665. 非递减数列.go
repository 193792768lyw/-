package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(checkPossibility([]int{-1, 4, 2, 3}))
}

func checkPossibility(nums []int) bool {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if i == 1 || nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
			count++
		}
	}
	return count <= 1
}


func checkPossibility(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			nums[i] = y
			if sort.IntsAreSorted(nums) {
				return true
			}
			nums[i] = x // 复原
			nums[i+1] = x
			return sort.IntsAreSorted(nums)
		}
	}
	return true
}

func checkPossibility(nums []int) bool {
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			cnt++
			if cnt > 1 {
				return false
			}
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
		}
	}
	return true
}
 /*
 例①： 4, 2, 5
 例②： 1, 4, 2, 5
 例③： 3, 4, 2, 5
  */