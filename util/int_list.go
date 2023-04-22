package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func List2Slice(list *ListNode) []int {
	var s []int
	for list != nil {
		s = append(s, list.Val)
		list = list.Next
	}

	return s
}

func Int2List(ints []int) *ListNode {
	dummy := &ListNode{}
	var p = dummy

	for _, v := range ints {
		p.Next = &ListNode{v, nil}
		p = p.Next
	}

	return dummy.Next
}

func Ints2Lists(ints [][]int) []*ListNode {
	var ret []*ListNode

	for _, n := range ints {
		ret = append(ret, Int2List(n))
	}

	return ret
}
