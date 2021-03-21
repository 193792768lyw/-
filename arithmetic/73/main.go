package main

import "fmt"

// 73. 矩阵置零
func main() {
	arr := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(arr)
	fmt.Println(arr)
}

type pos struct {
	col, row int
}

func setZeroes(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])
	mapZero := map[*pos]struct{}{}
	for i, arr := range matrix {
		for j, num := range arr {
			if num == 0 {
				mapZero[&pos{
					col: j,
					row: i,
				}] = struct{}{}
			}
		}
	}
	for pos, _ := range mapZero {
		key := pos.row
		value := pos.col
		for i := 0; i < col; i++ {
			matrix[key][i] = 0
		}
		for i := 0; i < row; i++ {
			matrix[i][value] = 0
		}
	}

}
