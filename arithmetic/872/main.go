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

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	res := make([]int, 0)

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			res = append(res, root.Val)
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root1)
	res2 := make([]int, len(res))
	copy(res2, res)
	res = make([]int, 0)

	dfs(root2)
	if len(res) != len(res2) {
		return false
	}
	for i := 0; i < len(res); i++ {

		if res[i] != res2[i] {
			return false
		}
	}
	return true
}
