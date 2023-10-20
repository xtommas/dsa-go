package main

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := LinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}

	list.append(5)
	list.append(7)
	list.append(9)

	expect("get", list.get(2), 9, t)
	expect("removeAt", list.removeAt(1), 7, t)
	expect("length", list.length, 2, t)

	list.append(11)

	expect("removeAt", list.removeAt(1), 9, t)
	expect("removeAt", list.remove(9), -1, t)
	expect("removeAt", list.removeAt(0), 5, t)
	expect("removeAt", list.removeAt(0), 11, t)
	expect("length", list.length, 0, t)

	list.prepend(5)
	list.prepend(7)
	list.prepend(9)
	list.insertAt(10, 2)

	expect("get", list.get(2), 10, t)
	expect("get", list.get(0), 9, t)
	expect("remove", list.remove(9), 9, t)
	expect("length", list.length, 3, t)
	expect("get", list.get(0), 7, t)

}
