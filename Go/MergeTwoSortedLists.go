package gosoln

import "github.com/leetcode/types"

func MergeTwoLists(l1 *types.ListNode, l2 *types.ListNode) *types.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = MergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoLists(l1, l2.Next)
		return l2
	}
}
