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

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	dummyHead := &ListNode{0, head}
	var lastTail *ListNode
	var kList *LinkedList

	lastTail = dummyHead

	for index, node := 0, dummyHead; node != nil; index, node = index+1, node.Next {
		if index%k == 1 {
			knode := &ListNode{node.Val, nil}
			kList = newLinkedList(knode)
			lastTail.Next = node
		} else {
			kList = kList.Prepend(node.Val)
		}

		if index > 0 && index%k == 0 {
			lastTail.Next = kList.Head
			lastTail = kList.Tail
		}
	}

	return dummyHead.Next
}

func (l *LinkedList) Prepend(val int) *LinkedList {
	if l == nil {
		node := &ListNode{val, nil}
		return newLinkedList(node)
	}
	l.Head = &ListNode{val, l.Head}
	return l
}

func newLinkedList(node *ListNode) *LinkedList {
	return &LinkedList{node, node}
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
