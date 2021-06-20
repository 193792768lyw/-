package main

import "fmt"

func main() {
	fmt.Println(SplitArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3))
}

func SplitArray(array []int, part int) [][]int {
	res := make([][]int, 0)

	length := len(array) / part //每一部分的长度
	less := len(array) % part

	tar := make([]int, 0)
	for i := 0; i < part-less; i++ {
		tar = array[:length]
		array = array[length:]
		res = append(res, tar)
	}
	for i := 0; i < less; i++ {
		tar = array[:length+1]
		array = array[length+1:]
		res = append(res, tar)
	}

	return res
}
