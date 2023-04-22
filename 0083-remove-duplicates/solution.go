package _083_remove_duplicates

import . "github.com/driventokill/leetcode-problems/util"

// DeleteDuplicates delete duplicates from list and return the new list
func DeleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	for fast != nil {
		if slow.Val == fast.Val {
			fast = fast.Next
		}
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
			fast = fast.Next
		}
	}
	slow.Next = nil

	return dummy.Next
}
