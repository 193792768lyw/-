package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*Node, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return res

}

func preorder1(root *Node) []int {

	res := make([]int, 0)
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for i := range node.Children {
			dfs(node.Children[i])
		}
	}

	dfs(root)
	return res

}
