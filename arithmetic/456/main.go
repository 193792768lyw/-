package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(find132pattern([]int{1, 2, 3, 4}))
}

// 方法一：使用暴力维护 3
func find132pattern1(nums []int) bool {
	len := len(nums)
	min := nums[0]
	for i := 1; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if min < nums[j] && nums[j] < nums[i] {
				return true
			}
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	return false
}

func min(one, two int) int {
	if one > two {
		return two
	}
	return one
}

/*
求任何位置的左边最小的元素 nums[i] ，可以提前遍历一次而得到；
使用「单调递减栈」，把 nums[j]  入栈时，需要把栈里面比它小的元素全都 pop 出来，由于越往栈底越大，所以 pop 出的最后一个元素，就是比 3 小的最大元素 nums[k] 。
判断如果 nums[i] < nums[k] ，那就说明得到了一个 132 模式。
*/

func find132pattern(nums []int) bool {
	N := len(nums)
	leftMin := make([]int, len(nums))
	leftMin[0] = math.MaxInt64
	for i := 1; i < N; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i-1])
	}

	stack := make([]int, 0)
	stack = append(stack, nums[N-1])
	for j := N - 2; j > 0; j-- {
		k := math.MinInt64
		for len(stack) != 0 && stack[len(stack)-1] < nums[j] {
			k = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if leftMin[j] < k {
			return true
		}
		stack = append(stack, nums[j])
	}

	return false

}
