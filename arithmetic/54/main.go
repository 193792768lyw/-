package main

import "fmt"

func main() {

	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func spiralOrder(matrix [][]int) []int {

	low, high, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	res := make([]int, 0)

	for low <= high && left <= right {
		// 上
		for i := left; i <= right; i++ {
			res = append(res, matrix[low][i])
		}
		low++
		if low > high {
			break
		}
		// 右
		for i := low; i <= high; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		// 下
		for i := right; i >= left; i-- {
			res = append(res, matrix[high][i])
		}
		high--
		if high < low {
			break
		}
		// 左
		for i := high; i >= low; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res
}
