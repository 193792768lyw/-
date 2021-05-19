package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 3))
}

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	results := make([]int, 0, m*n)
	pre := make([][]int, m+1)
	pre[0] = make([]int, n+1)
	for i, row := range matrix {
		pre[i+1] = make([]int, n+1)
		for j, val := range row {
			pre[i+1][j+1] = pre[i+1][j] ^ pre[i][j+1] ^ pre[i][j] ^ val
			results = append(results, pre[i+1][j+1])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(results)))
	return results[k-1]
}

//func kthLargestValue(matrix [][]int, k int) int {
//
//	for i := 0; i < len(matrix); i++ {
//		for j := 0; j < len(matrix[0]); j++ {
//			if j == 0 {
//
//				continue
//			}
//			matrix[i][j] = matrix[i][j-1] ^ matrix[i][j]
//		}
//	}
//	res := make([][]int, len(matrix))
//	for i := range res {
//		res[i] = make([]int, len(matrix[0]))
//	}
//	for i := 0; i < len(matrix); i++ {
//		if i == 0 {
//			res[0] = matrix[0]
//		}
//		for j := 0; j < len(matrix[0]); j++ {
//			t := 0
//			for k := 0; k <= i; k++ {
//				t ^= matrix[k][j]
//			}
//			res[i][j] = t
//		}
//	}
//	tt := make([]int, 0)
//	for i := range res {
//		for k := range res[i] {
//			tt = append(tt, res[i][k])
//		}
//	}
//	sort.Sort(sort.Reverse(sort.IntSlice(tt)))
//	return tt[k-1]
//}
