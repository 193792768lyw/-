package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 32))
}
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func combinationSum42(nums []int, target int) int {
	sort.Ints(nums)
	res := 0 // 181997601
	var dfs func(target int)
	dfs = func(target int) {
		if target == 0 {
			res++
			return
		}
		for i := 0; i < len(nums); i++ {
			if target-nums[i] < 0 {
				break
			}
			dfs(target - nums[i])
		}
	}
	dfs(target)
	return res

}
func combinationSum41(nums []int, target int) int {
	//sort.Ints(nums)
	ans := [][]int{}
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		for i := idx; i < len(nums); i++ {
			if target-nums[i] < 0 {
				break
			}
			comb = append(comb, nums[i])
			dfs(target-nums[i], 0)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	fmt.Println(ans)
	return len(ans)

}
