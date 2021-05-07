package main

import "fmt"

func main() {
	fmt.Println(decode([]int{1, 2, 3}, 1))
}

/*
输入：encoded = [1,2,3], first = 1
输出：[1,0,2,1]
解释：若 arr = [1,0,2,1] ，那么 first = 1 且 encoded = [1 XOR 0, 0 XOR 2, 2 XOR 1] = [1,2,3]


*/
func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first
	for k := range encoded {
		res[k+1] = encoded[k] ^ res[k]
	}
	return res
}
