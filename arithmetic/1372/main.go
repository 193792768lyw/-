package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	var dfs func(root *TreeNode, direction int) int
	dfs = func(root *TreeNode, direction int) int {
		if root == nil {
			return 0
		}
		if direction == 1 {
			return dfs(root.Left, 2)
		} else {
			return dfs(root.Right, 1)
		}
	}

	return max(dfs(root, 1), dfs(root, 2))

}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
