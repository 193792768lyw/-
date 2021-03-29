package main

// 92. 反转链表 II
func main() {

}

/**
 * Definition for singly-linked list.
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	newHead := new(ListNode)
	newHead.Next = head
	pre := newHead      // 前一个节点
	cur := newHead.Next // 当前遍历的节点
	count := 1
	firstNode := new(ListNode)

	for cur != nil {
		if count >= left {
			if count == left { // 记录反转的初始节点
				firstNode = cur
			}
			for count <= right { // 此处采用前插法完成翻转
				temp := cur
				cur = cur.Next
				temp.Next = pre.Next
				pre.Next = temp
				count++
			}
		}
		if count > right { // 当翻转结束后，直接用反转的初始节点连接后面的节点，然后就完成翻转
			firstNode.Next = cur
			break
		}
		pre = cur
		cur = cur.Next
		count++
	}

	return newHead.Next
}
