package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

/*
输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
*/
func subsetsWithDup1(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)

outer:
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
					continue outer
				}
				set = append(set, v)

			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return ans
}
func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	t := []int{}
	var dfs func(bool, int)
	dfs = func(choosePre bool, cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		dfs(false, cur+1)
		if !choosePre && cur > 0 && nums[cur-1] == nums[cur] {
			return
		}
		t = append(t, nums[cur])
		dfs(true, cur+1)
		t = t[:len(t)-1]
	}
	dfs(false, 0)
	return
}
