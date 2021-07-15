package leetcode

func removeNthNodeFromLastFastNode(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var fast, slow *ListNode
	prev := head
	for fast = head; fast != nil && n > 0; fast = fast.Next {
		n--
	}
	if fast == nil {
		return head.Next
	}
	for slow = head; fast != nil; slow = slow.Next {
		prev = slow.Next
		fast = fast.Next
	}
	prev.Next = slow.Next.Next
	return head
}