package main

import "fmt"

func main() {
	//fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 88))
	fmt.Println(searchMatrix([][]int{{1}}, 1))
}
func searchMatrix(matrix [][]int, target int) bool {
	index := 0
	for i := 0; i < len(matrix); i++ {
		if target >= matrix[i][0] && target <= matrix[i][len(matrix[0])-1] {
			index = i
			break
		}
	}
	return binarySearch(matrix[index], target)
}

func binarySearch(arr []int, target int) bool {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return true
		}
		if target < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
