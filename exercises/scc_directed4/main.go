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

func AddEdge(a, b *Node) {
	a.out = append(a.out, b)
	b.in = append(b.in, a)
}

func ReverseTopoSortHelper(node *Node, visited map[*Node]bool, sorted []*Node, nextIndex int) int {
	if visited[node] {
		return nextIndex
	}

	visited[node] = true

	for _, n := range node.in {
		nextIndex = ReverseTopoSortHelper(n, visited, sorted, nextIndex)
	}
	sorted[nextIndex] = node

	return nextIndex - 1
}

func ReverseTopSort(nodes []*Node) []*Node {
	// insertion index starts at len(nodes) - 1
	// DFS incoming edges
	// Return the next insertion index
	sorted := make([]*Node, len(nodes))
	visited := make(map[*Node]bool)
	nextIndex := len(nodes) - 1

	for _, n := range nodes {
		nextIndex = ReverseTopoSortHelper(n, visited, sorted, nextIndex)
		if nextIndex < 0 {
			break
		}
	}

	return sorted
}

func CalcSCC(n *Node, visited map[*Node]bool, components map[int][]*Node, currentComponent int) {
	if visited[n] {
		return
	}

	visited[n] = true

	for _, child := range n.out {
		CalcSCC(child, visited, components, currentComponent)
	}

	components[currentComponent] = append(components[currentComponent], n)
}

func CalcSCCs(sorted []*Node) map[int][]*Node {
	currentComponent := 0
	visited := make(map[*Node]bool, len(sorted))
	components := make(map[int][]*Node)

	for _, n := range sorted {
		if visited[n] {
			continue
		}

		currentComponent++
		CalcSCC(n, visited, components, currentComponent)
	}

	return components
}

func SCCs(nodes []*Node) map[int][]*Node {
	sorted := ReverseTopSort(nodes)
	return CalcSCCs(sorted)
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
