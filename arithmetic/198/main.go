package main

import (
	"fmt"
)

// 198. 打家劫舍
func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

/*
输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。

*/
func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[1] = nums[i]
		} else {
			dp[i+1] = max(dp[i], dp[i-1]+nums[i])
		}
	}
	res := 0
	for _, v := range dp {
		if res < v {
			res = v
		}
	}
	return res
}

func max(arr2 int, arr1 int) int {
	if arr1 > arr2 {
		return arr1
	}
	return arr2
}
