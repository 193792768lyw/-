package main

import (
	"fmt"
)

func main() {
	fmt.Println(xorQueries([]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}))
}

/*
输入：arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]
输出：[2,7,14,8]
解释：
数组中元素的二进制表示形式是：
1 = 0001
3 = 0011
4 = 0100
8 = 1000
查询的 XOR 值为：
[0,1] = 1 xor 3 = 2
[1,2] = 3 xor 4 = 7
[0,3] = 1 xor 3 xor 4 xor 8 = 14
[3,3] = 8
*/

func xorQueries(arr []int, queries [][]int) []int {

	arrXor := make([]int, len(arr))
	arrXor[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		arrXor[i] = arrXor[i-1] ^ arr[i]
	}
	res := make([]int, len(queries))
	for i := range queries {
		left, right := queries[i][0], queries[i][1]
		if left == 0 {
			res[i] = arrXor[right]
		} else if left == right {
			res[i] = arr[left]
		} else {
			res[i] = arrXor[right] ^ arrXor[left-1]

		}

	}
	return res
}
