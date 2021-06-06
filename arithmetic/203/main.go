package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	var f int64 = 1
	fmt.Println(time.Now().Add(time.Duration(f) * time.Hour))
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

func removeElements1(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
