package main

import "fmt"

// 给定一个数组，判断其有没有可能是一棵二叉查找树的前序遍历结果
func IsBST(arr []int) bool {

	var dfs func([]int) bool
	dfs = func(arr []int) bool {
		left := make([]int, 0)
		right := make([]int, 0)
		root := -1
		if len(arr) > 0 {
			root = arr[0]
			arr = arr[1:]
		} else {
			return true
		}
		index := 0
		for i, v := range arr {
			if root < v {
				left = append(left, v)
				index = i
			} else {
				break
			}
		}

		for i := index; i < len(arr); i++ {
			if root < arr[i] {
				right = append(right, arr[i])
			} else {
				return false
			}
		}
		return dfs(left) && dfs(right)

	}

	return dfs(arr)
}

func main() {
	arr := []int{3, 2, 1, 4}
	fmt.Println(IsBST(arr))
	arr = []int{4, 2, 1, 5, 3}
}

//
//3 | 2 1 | 4 5
//
//2 | 1
//4 | 5
