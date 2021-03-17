package main

import "fmt"

func main() {
	fmt.Println(numDistinct("babgbag", "bag"))
}

/*
动态规划
dp[i][j] 代表 T 前 i 字符串可以由 S j 字符串组成最多个数.
所以动态方程:
当 S[j] == T[i] , dp[i][j] = dp[i-1][j-1] + dp[i][j-1];
当 S[j] != T[i] , dp[i][j] = dp[i][j-1]
*/

/*
输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
(上箭头符号 ^ 表示选取的字母)
rabbbit
^^^^ ^^
rabbbit
^^ ^^^^
rabbbit
^^^ ^^^


输入：s = "babgbag", t = "bag"
输出：5
解释：
如下图所示, 有 5 种可以从 s 中得到 "bag" 的方案。
(上箭头符号 ^ 表示选取的字母)
babgbag
^^ ^
babgbag
^^    ^
babgbag
^    ^^
babgbag
  ^  ^^
babgbag
    ^^^

*/
func numDistinct(s string, t string) int {
	dp := make([][]int, len(t)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(s)+1)
	}

	for j := 0; j < len(s)+1; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < len(t)+1; i++ {
		for j := 1; j < len(s)+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]

}
