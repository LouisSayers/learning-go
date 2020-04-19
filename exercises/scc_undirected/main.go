package main

import "fmt"

type Node struct {
	val string
	out []*Node
}

func (n *Node) String() string {
	return n.val
}

func NewNode(val string) *Node {
	return &Node{val, []*Node{}}
}

func AddEdge(n1, n2 *Node) {
	n1.out = append(n1.out, n2)
	n2.out = append(n2.out, n1)
}

func DFS(n *Node, componentMap map[*Node]int, cc int) {
	if componentMap[n] != 0 {
		return
	}

	componentMap[n] = cc
	for _, child := range n.out {
		DFS(child, componentMap, cc)
	}
}

func DFSIterative(n *Node, componentMap map[*Node]int, cc int) {
	stack := []*Node{n}

	for len(stack) != 0 {
		lastIndex := len(stack) - 1
		nextItem := stack[lastIndex]
		stack = stack[:lastIndex]

		// skip if the node is mapped to a connected component
		if componentMap[nextItem] != 0 {
			continue
		}

		// otherwise, add this item to the cc, process the out nodes
		componentMap[nextItem] = cc
		for _, child := range nextItem.out {
			stack = append(stack, child)
		}
	}
}

func PrintCC(arr []*Node) {
	componentMap := make(map[*Node]int)
	var cc int

	for _, node := range arr {
		if componentMap[node] != 0 {
			continue
		}

		cc++
		DFSIterative(node, componentMap, cc)
	}

	ccs := make(map[int][]*Node)
	for n, cc := range componentMap {
		ccs[cc] = append(ccs[cc], n)
	}

	fmt.Printf("Result: %v\n", ccs)
}

func main() {
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")
	f := NewNode("f")
	g := NewNode("g")

	// CC 1
	AddEdge(a, b)
	AddEdge(b, c)
	AddEdge(b, g)
	AddEdge(c, g)

	// CC 2
	AddEdge(d, e)
	AddEdge(d, f)
	AddEdge(e, f)

	allNodes := []*Node{a, b, c, d, e, f, g}
	PrintCC(allNodes)
}
