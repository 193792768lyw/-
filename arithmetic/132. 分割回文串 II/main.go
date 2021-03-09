package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCut("aab"))
}
func minCut(s string) int {
	l := len(s)
	res := math.MaxInt32
	if l == 0 {
		return res
	}
	//path := make([]string, 0)
	dp := longestPalindrome(s)

	var dfs func(int, int)
	dfs = func(index int, deep int) {
		if index == l {
			//temp := make([]string, len(path))
			//copy(temp, path)
			//res = append(res, temp)
			if deep < res {
				res = deep
			}
			return
		}

		for i := index; i < l; i++ {
			if dp[index][i] > 0 {
				//如果前缀字符串是回文，则可以产生分支和结点；
				//path = append(path, s[index:i+1])
				dfs(i+1, deep+1)
				//path = path[:len(path)-1]
			}
		}
	}

	dfs(0, 0)
	return res
}

func longestPalindrome(s string) [][]int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	// 状态转移方程：在 s[i] == s[j] 的时候，dp[i][j] 参考 dp[i + 1][j - 1]
	for right := 0; right < n; right++ {
		// 注意：left <= right 取等号表示 1 个字符的时候也需要判断
		for left := 0; left <= right; left++ {
			if s[left] == s[right] && (right-left <= 2 || dp[left+1][right-1] > 0) {
				dp[left][right] = 1
			}
		}
	}
	return dp
}
