package main

func main() {

}

/**
 * Definition for singly-linked list. */
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
  输入：head = [1,1,2,3,3]
  输出：[1,2,3]
*/
func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	end := head

	for head != nil {
		for head != nil && head.Val == end.Val {
			head = head.Next
		}
		if head == nil {
			break
		}
		if head.Val != end.Val {
			end.Next = head
			end = head
		}
		head = head.Next
	}
	if end != nil {
		end.Next = nil

	}
	return newHead.Next
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}
