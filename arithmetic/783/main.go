package main

import "math"

func main() {

}

/**
 * Definition for a binary tree node. */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	res := math.MaxInt64
	nums := make([]int, 0)
	var dd func(root *TreeNode)
	dd = func(root *TreeNode) {
		if root == nil {
			return
		}
		dd(root.Left)
		if len(nums) == 0 {
			nums = append(nums, root.Val)
		} else {
			if res > root.Val-nums[0] {
				res = root.Val - nums[0]
				nums[0] = root.Val
			} else {
				nums[0] = root.Val
			}
		}
		dd(root.Right)
	}
	dd(root)
	return res

}
