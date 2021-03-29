package main

func main() {

}

/**
 * Definition for singly-linked list. */
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	elements := make([]int, 0)
	for head != nil {
		elements = append(elements, head.Val)
		head = head.Next
	}

	mv := k % len(elements)
	reversal(elements, 0, len(elements)-mv-1)
	reversal(elements, len(elements)-mv, len(elements)-1)
	reversal(elements, 0, len(elements)-1)
	newHead := new(ListNode)
	h := newHead
	for i := range elements {
		h.Val = elements[i]
		if i != len(elements)-1 {
			h.Next = new(ListNode)
			h = h.Next
		}
	}
	return newHead
}

func reversal(arr []int, left, right int) {
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
