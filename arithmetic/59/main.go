package main

import "fmt"

// 59. 螺旋矩阵 II
func main() {
	fmt.Println(generateMatrix(1))
}
func generateMatrix(n int) [][]int {

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	low, high, left, right := 0, n-1, 0, n-1
	num := 1

	for low <= high && left <= right {
		// 上
		for i := left; i <= right; i++ {
			res[low][i] = num
			num++
		}
		low++
		if low > high {
			break
		}
		// 右
		for i := low; i <= high; i++ {
			res[i][right] = num
			num++
		}
		right--
		if right < left {
			break
		}
		// 下
		for i := right; i >= left; i-- {
			res[high][i] = num
			num++
		}
		high--
		if high < low {
			break
		}
		// 左
		for i := high; i >= low; i-- {
			res[i][left] = num
			num++
		}
		left++
	}
	return res
}
