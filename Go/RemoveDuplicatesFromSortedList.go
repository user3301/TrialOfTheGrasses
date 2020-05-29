package gosoln

import "github.com/leetcode/types"

func DeleteDuplicates(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head.Next = DeleteDuplicates(head.Next)
	if head.Val == head.Next.Val {
		head = head.Next
	}
	return head
}
