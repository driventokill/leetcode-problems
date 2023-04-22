package _019

import (
	. "github.com/driventokill/leetcode-problems/util"
)

func RemoveNthFromEnd(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	n := FindFromEnd(head, k+1)
	n.Next = n.Next.Next
	return dummy.Next
}

func FindFromEnd(list *ListNode, k int) *ListNode {
	p1 := list
	// find the kth node from head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}

	p2 := list
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}
