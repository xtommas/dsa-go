package main

import (
	"errors"
	"fmt"
)

type DoublyLinkedList struct {
	length int
	head   *DoubleNode
	tail   *DoubleNode
}

func (l *DoublyLinkedList) prepend(item int) {
	node := DoubleNode{
		value: item,
		next:  nil,
		prev:  nil,
	}

	if l.length == 0 {
		l.head = &node
		l.tail = &node
		l.length++
		return
	}
	node.next = l.head
	l.head.prev = &node
	l.head = &node
	l.length++
}

func (l *DoublyLinkedList) insertAt(item, idx int) {
	if idx > l.length {
		fmt.Println(errors.New("Index out of bounds"))
	} else if idx == l.length {
		l.append(item)
		return
	} else if idx == 0 {
		l.prepend(item)
		return
	}

	curr := l.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}

	node := DoubleNode{
		value: item,
		next:  nil,
		prev:  nil,
	}

	node.next = curr
	node.prev = curr.prev
	curr.prev.next = &node
	curr.prev = &node
	l.length++
}

func (l *DoublyLinkedList) append(item int) {
	node := DoubleNode{
		value: item,
		next:  nil,
		prev:  nil,
	}

	if l.length == 0 || l.head == nil || l.tail == nil {
		l.head = &node
		l.tail = &node
		l.length++
		return
	}

	node.prev = l.tail
	l.tail.next = &node
	l.tail = &node
	l.length++
}

func (l *DoublyLinkedList) remove(item int) int {
	if l.length == 0 {
		return -1
	}

	curr := l.head
	for i := 0; i < l.length; i++ {
		if curr.value == item {
			break
		}
		curr = curr.next
	}

	// item not found
	if curr == nil {
		return -1
	}

	// list length 1
	if l.length == 1 {
		out := l.head.value
		l.head = nil
		l.tail = nil
		l.length--
		return out
	}

	// update the pointer for removal
	if curr.prev != nil {
		curr.prev.next = curr.next
	}

	if curr.next != nil {
		curr.next.prev = curr.prev
	}

	if curr == l.head {
		l.head = curr.next
	}

	if curr == l.tail {
		l.tail = curr.prev
	}

	curr.prev = nil
	curr.next = nil
	l.length--
	return curr.value
}

func (l *DoublyLinkedList) get(idx int) int {
	if idx > l.length || idx < 0 {
		return -1
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

func (l *DoublyLinkedList) removeAt(idx int) int {
	if idx > l.length || idx < 0 {
		return -1
	}
	if l.length == 0 {
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
		l.head = l.head.next
		l.head.prev.next = nil
		l.head.prev = nil
		l.length--
		return val
	}
	if idx == l.length-1 {
		val := l.tail.value
		l.tail = l.tail.prev
		l.tail.next.prev = nil
		l.tail.next = nil
		l.length--
		return val
	}
	curr := l.head
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.next
	}
	curr.prev.next = curr.next
	curr.next.prev = curr.prev
	curr.next = nil
	curr.prev = nil
	l.length--
	return curr.value
}
