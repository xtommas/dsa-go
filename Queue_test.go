package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Queue{
		head:   nil,
		tail:   nil,
		length: 0,
	}

	queue.enqueue(5)
	queue.enqueue(7)
	queue.enqueue(9)

	expect("dequeue", queue.dequeue(), 5, t)
	expect("length", queue.length, 2, t)

	queue.enqueue(11)

	expect("dequeue", queue.dequeue(), 7, t)
	expect("dequeue", queue.dequeue(), 9, t)
	expect("peek", queue.peek(), 11, t)
	expect("dequeue", queue.dequeue(), 11, t)
	expect("dequeue", queue.dequeue(), -1, t)
	expect("length", queue.length, 0, t)

	queue.enqueue(69)
	expect("peek", queue.peek(), 69, t)
	expect("length", queue.length, 1, t)
}
