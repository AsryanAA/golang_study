package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func New(initialValue int) *Node {
	return &Node{
		value: initialValue,
	}
}

func (n *Node) Append(i int) {
	current := n

	for {
		if current.next == nil {
			newNode := New(i)
			current.next = newNode

			return
		}

		current = current.next
	}
}

func main() {
	node := New(16)

	node.Append(42)
	node.Append(22)

	for {
		fmt.Println(node.value)
		node = node.next

		if node == nil {
			return
		}
	}
}