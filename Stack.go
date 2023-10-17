package main

type Stack struct {
	length int
	top    *Node
}

func (s *Stack) push(item int) {
	newNode := Node{
		value: item,
		next:  nil,
	}
	aux := s.top
	s.top = &newNode
	s.top.next = aux
	s.length++
}

func (s *Stack) pop() int {
	if s.length == 0 {
		return -1
	} else {
		aux := s.top
		s.top = s.top.next
		aux.next = nil
		s.length--
		return aux.value
	}
}

func (s *Stack) peek() int {
	return s.top.value
}
