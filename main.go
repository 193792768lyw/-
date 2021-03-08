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

	// f(i) 代表考虑前 i 个字符的最小分割次数
	f := make([]int, l)
	for j := 1; j < l; j++ {

		// 如果 [0,j] 这一段直接构成回文，则无须分割
		if dp[0][j] > 0 {
			f[j] = 0
			// 如果无法直接构成回文
			// 那么对于第 j 个字符，有使用分割次数，或者不使用分割次数两种选择
		} else {
			// 下边两种决策也能够合到一个循环当中去做，但是需要先将 f[i] 预设为一个足够大的数，因此干脆拆开来做
			// 独立使用一次分割次数
			//f[j] = f[j-1] + 1
			f[j] = math.MaxInt64
			// 第 j 个字符本身不独立使用分割次数
			// 代表要与前面的某个位置 i 形成区间 [i,j]，使得 [i, j] 形成回文，[i, j] 整体消耗一次分割次数
			for i := 1; i <= j; i++ {
				if dp[i][j] > 0 {
					f[j] = min(f[j], f[i-1]+1)
				}
			}
		}
	}
	return f[l-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
