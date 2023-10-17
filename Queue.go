package main

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (q *Queue) enqueue(item int) {
	newNode := Node{
		value: item,
		next:  nil,
	}

	if q.length == 0 {
		q.head = &newNode
		q.tail = &newNode
	} else {
		q.tail.next = &newNode
		q.tail = &newNode
	}
	q.length++
}

func (q *Queue) dequeue() int {
	if q.length == 0 {
		return -1
	} else {
		auxNode := q.head
		q.head = q.head.next
		auxNode.next = nil
		q.length--
		return auxNode.value
	}

}

func (q *Queue) peek() int {
	return q.head.value
}
