package main

type Node struct {
	value int
	next  *Node
}

type DoubleNode struct {
	value int
	next  *DoubleNode
	prev  *DoubleNode
}
