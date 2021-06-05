package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseFloat("7777", 10))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var len1, len2 int
	new1 := headA
	new2 := headB
	for headA != nil {
		len1++
		headA = headA.Next
	}
	for headB != nil {
		len2++
		headB = headB.Next
	}
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			new1 = new1.Next
		}
	} else {
		for i := 0; i < len2-len1; i++ {
			new2 = new2.Next
		}
	}
	for new1 != nil {
		if new1 == new2 {
			return new1
		} else {
			new1 = new1.Next
			new2 = new2.Next
		}
	}
	return nil
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
