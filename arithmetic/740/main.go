package main

import "sort"

func main() {

}

/*
输入：nums = [3,4,2]
输出：6
解释：
删除 4 获得 4 个点数，因此 3 也被删除。
之后，删除 2 获得 2 个点数。总共获得 6 个点数。

*/
//func deleteAndEarn(nums []int) int {
//	maxVal := 0
//	for _, val := range nums {
//		maxVal = max(maxVal, val)
//	}
//	sum := make([]int, maxVal+1)
//	for _, val := range nums {
//		sum[val] += val
//	}
//	return rob(sum)
//}
//
//func rob(nums []int) int {
//	first, second := nums[0], max(nums[0], nums[1])
//	for i := 2; i < len(nums); i++ {
//		first, second = second, max(first+nums[i], second)
//	}
//	return second
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func deleteAndEarn(nums []int) (ans int) {
	sort.Ints(nums)
	sum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if val := nums[i]; val == nums[i-1] {
			sum[len(sum)-1] += val
		} else if val == nums[i-1]+1 {
			sum = append(sum, val)
		} else {
			ans += rob(sum)
			sum = []int{val}
		}
	}
	ans += rob(sum)
	return
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
