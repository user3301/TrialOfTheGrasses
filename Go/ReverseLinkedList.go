package gosoln

import "github.com/leetcode/types"

func ReverseList(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
