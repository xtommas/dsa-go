package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := Stack{
		length: 0,
		top:    nil,
	}

	stack.push(5)
	stack.push(7)
	stack.push(9)

	expect("pop", stack.pop(), 9, t)
	expect("length", stack.length, 2, t)

	stack.push(11)
	expect("pop", stack.pop(), 11, t)
	expect("pop", stack.pop(), 7, t)
	expect("peek", stack.peek(), 5, t)
	expect("pop", stack.pop(), 5, t)
	expect("pop", stack.pop(), -1, t)

	stack.push(69)
	expect("peek", stack.peek(), 69, t)
	expect("length", stack.length, 1, t)
}
