package main

import "fmt"

func main() {
	n := Constructor22([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	fmt.Println(n.SumRegion(1, 1, 2, 2))
}

type NumMatrix struct {
	matrix    [][]int
	preMatrix [][]int
}

func Constructor22(matrix [][]int) NumMatrix {
	preSum := make([][]int, len(matrix))
	for i, v := range matrix {
		sum := make([]int, len(v))
		sum[0] = v[0]
		for j := 1; j < len(v); j++ {
			sum[j] = sum[j-1] + v[j]
		}
		preSum[i] = sum
	}

	return NumMatrix{
		matrix:    matrix,
		preMatrix: preSum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		if col1 == 0 {
			sum += this.preMatrix[i][col2]
		}else{
			sum += this.preMatrix[i][col2] - this.preMatrix[i][col1-1]

		}
	}
	return sum
}