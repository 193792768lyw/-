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

//type BSTIterator struct {
//	arr []int
//}
//
//func Constructor(root *TreeNode) (it BSTIterator) {
//	it.inorder(root)
//	return
//}
//
//func (it *BSTIterator) inorder(node *TreeNode) {
//	if node == nil {
//		return
//	}
//	it.inorder(node.Left)
//	it.arr = append(it.arr, node.Val)
//	it.inorder(node.Right)
//}
//
//func (it *BSTIterator) Next() int {
//	val := it.arr[0]
//	it.arr = it.arr[1:]
//	return val
//}
//
//func (it *BSTIterator) HasNext() bool {
//	return len(it.arr) > 0
//}

type BSTIterator struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{cur: root}
}

func (it *BSTIterator) Next() int {
	// 往左走到最底下
	for node := it.cur; node != nil; node = node.Left {
		it.stack = append(it.stack, node)
	}

	it.cur, it.stack = it.stack[len(it.stack)-1], it.stack[:len(it.stack)-1]

	val := it.cur.Val
	it.cur = it.cur.Right
	return val
}

func (it *BSTIterator) HasNext() bool {
	return it.cur != nil || len(it.stack) > 0
}
