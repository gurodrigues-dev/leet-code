package main

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
