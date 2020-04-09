package main

import "fmt"

type Node struct {
	next *Node
	val string
}

func (n *Node) Reverse() {
	n.reverse(nil)
}

func (n *Node) reverse(node *Node) {
	if n.next == nil {
		n.next = node
		return
	}

	n.next.reverse(n)
	n.next = node
}

func (n *Node) hasNext() bool {
	return n.next != nil
}

func reverseIterative(root *Node) {
	var prev *Node

	current := root

	for current != nil {
		temp := current.next
		current.next = prev
		prev = current
		current = temp
	}
}


func main() {
	n1 := &Node{nil, "c"}
	n2 := &Node{n1, "b"}
	n3 := &Node{n2, "a"}

	reverseIterative(n3)

	var current *Node
	for current = n1; current != nil; current = current.next {
		fmt.Printf("%v->", current.val)
	}
}
