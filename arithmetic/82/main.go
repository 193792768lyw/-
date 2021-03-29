package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	r := gjson.Parse("{\"level\":\"142\",\"honour\":\"33\",\"nickname\":\"狗子～你醒...\"}")
	fmt.Println(r.Get("nickname").Value().(string))
}

/**输入: 1->2->3->3->4->4->5
输出: 1->2->5
 * Definition for singly-linked list. */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{}
	end := newHead
	last := head

	for head != nil {
		head = head.Next
		if head == nil || head.Val != last.Val {
			end.Next = last
			end = last
		} else { // 1->1->1->2->3
			for head != nil && head.Val == last.Val {
				head = head.Next
			}
		}
		last = head

	}
	end.Next = nil
	return newHead.Next
}
