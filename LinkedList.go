package main

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func (l *LinkedList) prepend(item int) {
	node := Node{
		value: item,
		next:  nil,
	}

	if l.length == 0 {
		l.head = &node
		l.tail = &node
		l.length++
		return
	}

	node.next = l.head
	l.head = &node
	l.length++
}

func (l *LinkedList) insertAt(item, idx int) {
	if idx > l.length || idx < 0 {
		fmt.Println(errors.New("index out of bounds"))
	} else if idx == l.length {
		l.append(item)
		return
	} else if idx == 0 {
		l.prepend(item)
		return
	}

	curr := l.head
	prev := curr
	for i := 0; i < idx; i++ {
		prev = curr
		curr = curr.next
	}

	node := Node{
		value: item,
	}
	node.next = curr
	prev.next = &node
	l.length++
}

func (l *LinkedList) append(item int) {
	node := Node{
		value: item,
	}
	if l.length == 0 || l.head == nil || l.tail == nil {
		l.head = &node
		l.tail = &node
		l.length++
		return
	}

	l.tail.next = &node
	l.tail = &node
	l.length++
}

func (l *LinkedList) remove(item int) int {
	if l.length == 0 {
		return -1
	}

	curr := l.head
	prev := curr
	for i := 0; i < l.length; i++ {
		if curr.value == item {
			break
		}
		prev = curr
		curr = curr.next
	}

	if curr == nil {
		return -1
	}

	if l.length == 1 {
		out := l.head.value
		l.head = nil
		l.tail = nil
		l.length--
		return out
	}

	if prev != nil {
		prev.next = curr.next
	}

	if curr == l.head {
		l.head = curr.next
	}

	curr.next = nil
	l.length--
	return curr.value
}

func (l *LinkedList) get(idx int) int {
	if idx > l.length || idx < 0 {
		fmt.Println(errors.New("index out of bounds"))
	}

	if idx == 0 {
		return l.head.value
	}

	if idx == l.length-1 {
		return l.tail.value
	}

	curr := l.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}

	return curr.value
}

func (l *LinkedList) removeAt(idx int) int {
	if idx > l.length || idx < 0 {
		return -1
	}

	if l.head == nil {
		return -1
	}

	if l.length == 1 {
		val := l.head.value
		l.head = nil
		l.length--
		return val
	}

	if idx == 0 {
		val := l.head.value
		tmp := l.head
		l.head = l.head.next
		tmp.next = nil
		l.length--
		return val
	}

	curr := l.head
	prev := curr
	for i := 0; curr != nil && i < idx; i++ {
		prev = curr
		curr = curr.next
	}

	prev.next = curr.next
	curr.next = nil
	l.length--
	return curr.value
}
