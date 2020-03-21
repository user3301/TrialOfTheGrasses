package Go

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
