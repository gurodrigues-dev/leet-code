package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(value int) {
	newNode := &Node{
		value: value,
		next:  s.top,
	}
	s.top = newNode
}

func (s *Stack) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	val := s.top.value
	s.top = s.top.next
	return val, true
}

func (s *Stack) Peek() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	return s.top.value, true
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Print() {
	for node := s.top; node != nil; node = node.next {
		fmt.Printf("%d -> ", node.value)
	}
	fmt.Println("nil")
}
