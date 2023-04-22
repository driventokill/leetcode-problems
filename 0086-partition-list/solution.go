package partitionlist

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	smallHead := &ListNode{}
	smallTail := smallHead

	largeHead := &ListNode{}
	largeTail := largeHead

	node := head
	
	for node != nil {
		if node.Val < x {
			smallTail.Next = node
			smallTail = smallTail.Next
		} else {
			largeTail.Next = node
			largeTail = largeTail.Next
		}
		// 断开原链表中的 next 指针
		temp := node.Next
		node.Next = nil
		node = temp 
	}
	
	smallTail.Next = largeHead.Next

	return smallHead.Next
}
