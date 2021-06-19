package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(maxLength([]string{"bbb"}))
}

func maxLength(arr []string) (ans int) {
	masks := []int{}
outer:
	for _, s := range arr {
		mask := 0
		for _, ch := range s {
			ch -= 'a'
			if mask>>ch&1 > 0 { // 若 mask 已有 ch，则说明 s 含有重复字母，无法构成可行解
				continue outer
			}
			mask |= 1 << ch // 将 ch 加入 mask 中
		}
		masks = append(masks, mask)
	}

	var backtrack func(int, int)
	backtrack = func(pos, mask int) {
		if pos == len(masks) {
			ans = max(ans, bits.OnesCount(uint(mask)))
			return
		}
		if mask&masks[pos] == 0 { // mask 和 masks[pos] 无公共元素
			backtrack(pos+1, mask|masks[pos])
		}
		backtrack(pos+1, mask)
	}
	backtrack(0, 0)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
