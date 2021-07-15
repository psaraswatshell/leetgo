package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{0, nil}
	var curr = head
	var carry = 0
	for l1 != nil || l2 != nil {
		var sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}
	return head.Next
}