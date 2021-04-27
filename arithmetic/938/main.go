package main

func main() {

}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val >= low && node.Val <= high {
			res += node.Val
		}
		if node.Val < low {
			dfs(node.Right)
		} else if node.Val > high {
			dfs(node.Left)
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}

	}
	dfs(root)
	return res
}
