package main

import "fmt"

func main() {
	//输入：s = "AABABBA", k = 1
	//输出：4
	//解释：
	//将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
	//子串 "BBBB" 有最长重复字母, 答案为 4。

	fmt.Println(characterReplacement("AABBBCC", 0))
}
func characterReplacement(s string, k int) int {
	cnt := [26]int{}
	maxCnt, left := 0, 0
	for right, ch := range s {
		cnt[ch-'A']++
		maxCnt = max(maxCnt, cnt[ch-'A'])
		if right-left+1-maxCnt > k {
			cnt[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
