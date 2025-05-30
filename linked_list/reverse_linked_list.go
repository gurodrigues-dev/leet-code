package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkedList(head *ListNode) *ListNode {
	var newList *ListNode // start -> nil

	for head != nil {
		nextNode := head.Next
		head.Next = newList
		newList = head
		head = nextNode
	}
	return newList
}

func main() {
	head := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("Original: ")
	printList(head)

	reversed := reverseLinkedList(head)
	fmt.Print("Reversa: ")
	printList(reversed)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, v := range values[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}
