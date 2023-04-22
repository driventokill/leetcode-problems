package mergeksortedlists

import "container/heap"

import . "github.com/driventokill/leetcode-problems/util"

func mergeKListsByHeap(lists []*ListNode) *ListNode {
	pq := make(Lists, 0)
	heap.Init(&pq)

	// put the first nodes of sorted lists
	for _, l := range lists {
		if l != nil {
			heap.Push(&pq, l)
		}
	}

	dummy := &ListNode{}
	p := dummy

	// pop from heap and put the next list node one by one
	for pq.Len() > 0 {
		n := heap.Pop(&pq).(*ListNode)
		p.Next = n
		p = p.Next

		if n.Next != nil {
			heap.Push(&pq, n.Next)
			n.Next = nil
		}
	}
	return dummy.Next
}

type Lists []*ListNode

func (ls Lists) Len() int {
	return len(ls)
}

func (ls Lists) Swap(i, j int) {
	ls[i], ls[j] = ls[j], ls[i]
}

func (ls Lists) Less(i, j int) bool {
	return ls[i].Val <= ls[j].Val
}

func (ls *Lists) Push(v interface{}) {
	*ls = append(*ls, v.(*ListNode))
}

func (ls *Lists) Pop() interface{} {
	tmp := *ls
	l := tmp.Len()
	var res = tmp[l-1]
	*ls = tmp[:l-1]
	return res
}

// ---- merge brutally

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	dummyTail := dummy

	finished := false

	for !finished {
		finished = true
		minNode := &ListNode{int(^uint(0) >> 1), nil}
		minNodeIdx := -1

		// find the mininum node
		for i, list := range lists {
			if list == nil {
				continue
			}

			finished = false
			if minNode.Val > list.Val {
				minNode = list
				minNodeIdx = i
			}
		}

		// empty lists
		if minNodeIdx == -1 {
			continue
		}

		// move the min node to return list
		temp := minNode.Next
		minNode.Next = nil
		lists[minNodeIdx] = temp

		dummyTail.Next = minNode
		dummyTail = dummyTail.Next
	}

	return dummy.Next
}
