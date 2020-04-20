package main

import "fmt"

type Node struct {
	val string
	in []*Node
	out []*Node
}

func (n *Node) String() string {
	return n.val
}

func NewNode(val string) *Node {
	return &Node{val, []*Node{}, []*Node{}}
}

func AddEdge(n1, n2 *Node) {
	n1.out = append(n1.out, n2)
	n2.in = append(n2.in, n1)
}

func AllReachableFrom(n *Node) []*Node {
	visited := map[*Node]bool{ n: true }
	stack := []*Node{n}
	var nodes []*Node

	for len(stack) != 0 {
		lastItemIndex := len(stack) - 1
		item := stack[lastItemIndex]
		stack = stack[:lastItemIndex]

		for _, n := range item.out {
			if !visited[n] {
				nodes = append(nodes, n)
				stack = append(stack, n)
			}
			visited[n] = true
		}
	}

	return nodes
}

// O(n^2 + ne^2)
func SCCs(nodes []*Node) map[int][]*Node {
	sccs := make(map[int][]*Node)
	reachable := make(map[*Node][]*Node)

	for _, n := range nodes {
		reachable[n] = AllReachableFrom(n)
	}

	fmt.Println("Reachable: ", reachable)

	foundSCC := make(map[*Node]bool)
	currentSCC := 0
	for _, n := range nodes {
		if foundSCC[n] {
			continue
		}
		foundSCC[n] = true
		currentSCC++
		sccs[currentSCC] = append(sccs[currentSCC], n)
		for _, n2 := range reachable[n] {
			for _, n3 := range reachable[n2] {
				if n == n3 {
					foundSCC[n2] = true
					sccs[currentSCC] = append(sccs[currentSCC], n2)
				}
			}
		}
	}

	return sccs
}

func main() {
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")
	f := NewNode("f")
	g := NewNode("g")
	h := NewNode("h")
	i := NewNode("i")
	j := NewNode("j")

	AddEdge(a, b)
	AddEdge(b, g)
	AddEdge(g, c)
	AddEdge(c, b)
	AddEdge(c, e)
	AddEdge(e, f)
	AddEdge(f, d)
	AddEdge(d, e)
	AddEdge(e, h)
	AddEdge(h, i)
	AddEdge(i, j)
	AddEdge(j, h)

	nodes := []*Node{a, b, c, d, e, f, g, h, i, j}
	sccs := SCCs(nodes)
	fmt.Printf("SCCs: %v\n", sccs)
}
