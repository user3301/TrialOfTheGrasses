package gosoln

import "github.com/leetcode/types"

func HasCycle(head *types.ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head

	for fast.Next != nil && slow.Next != nil {
		if fast.Next.Next == nil {
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}
