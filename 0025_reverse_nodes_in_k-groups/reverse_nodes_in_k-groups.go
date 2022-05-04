package reversenodesinkgroups

import (
	"strconv"
	"strings"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	dummyHead := &ListNode{0, head}
	var lastTail, kHead *ListNode
	lastTail = dummyHead
	for index, node := 0, dummyHead; node != nil; index, node = index+1, node.Next {
		if index%k == 1 {
			kHead = &ListNode{node.Val, nil}
			lastTail.Next = node
		} else {
			kHead = &ListNode{node.Val, kHead}
		}

		if index > 0 && index%k == 0 {
			lastTail.Next = kHead
			lastTail = tail(kHead)
		}
	}

	return dummyHead.Next
}

func tail(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var tail *ListNode

	for node := head; node != nil; node = node.Next {
		tail = node
	}

	return tail
}

// ----- assist funcs -----
func newList(values ...int) *ListNode {
	dummyHead := &ListNode{0, nil}

	for head, i := dummyHead, 0; i < len(values); head, i = head.Next, i+1 {
		head.Next = &ListNode{values[i], nil}
	}

	return dummyHead.Next
}

func (head *ListNode) String() string {
	var values []string
	for node := head; node != nil; node = node.Next {
		values = append(values, strconv.Itoa(node.Val))
	}

	return strings.Join(values, "->")
}
