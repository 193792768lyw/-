package main

import "fmt"

// 213. 打家劫舍 II
func main() {
	fmt.Println(rob([]int{200, 3, 140, 20, 10}))
}

/*
输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。

*/

func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])

	}

	one := robF(nums[:len(nums)-1])
	two := robF(nums[1:])
	return max(one, two)

}

func robF(nums []int) int {
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
