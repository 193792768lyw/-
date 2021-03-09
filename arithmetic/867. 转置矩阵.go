package main

import "fmt"

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
func transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))

	for i := 0; i < len(matrix[0]); i++ {
		row := make([]int, len(matrix))
		for j := 0; j < len(matrix); j++ {
			row[j] = matrix[j][i]
		}
		res[i] = row
	}
	return res
}
