package main

import "fmt"

func main() {
	//fmt.Println(peakIndexInMountainArray([]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}))
	fmt.Println(levelOrder(nil))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	//if root == nil {
	//	return res
	//}
	quene := make([]*TreeNode, 0)
	quene = append(quene, root)
	//lev := 1
	for len(quene) > 0 {
		temp := make([]int, 0)
		quene1 := make([]*TreeNode, 0)
		for _, r := range quene {
			temp = append(temp, r.Val)
			if r.Left != nil {
				quene1 = append(quene1, r.Left)
			}
			if r.Right != nil {
				quene1 = append(quene1, r.Right)
			}
		}

		//if lev%2 == 0 {
		//	for l, r := 0, len(temp)-1; l < r; {
		//		temp[l], temp[r] = temp[r], temp[l]
		//		l++
		//		r--
		//	}
		//
		//}

		res = append(res, temp)
		//lev++
		quene = quene1
	}

	return res
}

func peakIndexInMountainArray(arr []int) int {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return i - 1
		}
	}
	return 0
}
