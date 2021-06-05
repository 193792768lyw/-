package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{}
	cur := newHead
	for head != nil {
		if head.Val != val {
			cur.Next = head
			cur = cur.Next
		}
		head = head.Next
	}
	cur.Next = nil

	return newHead.Next
}
