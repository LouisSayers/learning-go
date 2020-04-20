package main

import "fmt"

type Node struct {
	val string
	out []*Node
	in []*Node
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

func InverseSort(n *Node, visited map[*Node]bool, sorted []*Node, index int) int {
	if visited[n] {
		return index
	}

	stack := []*Node{n}
	visited[n] = true

	for len(stack) != 0 {
		current := stack[len(stack) - 1]

		parentsAllVisited := true
		for _, node := range current.in {
			if !visited[node] {
				stack = append(stack, node)
				parentsAllVisited = false
				visited[node] = true
			}
		}

		if parentsAllVisited {
			fmt.Printf("index: %d, current: %v, stack (%d): %v, visited: %v\n", index, current, len(stack), stack, visited)
			sorted[index] = current
			index--
			stack = stack[:(len(stack) - 1)]
			visited[current] = true
		}
	}

	return index
}

func InverseTopoSort(nodes, sorted []*Node) {
  visited := make(map[*Node]bool)
	nextInsertIndex := len(nodes) - 1

  for _, n := range nodes {
	  nextInsertIndex = InverseSort(n, visited, sorted, nextInsertIndex)
  }
}

func ConnectedComponents(sorted []*Node) map[int][]*Node {
	components := make(map[int][]*Node)
	visited := make(map[*Node]bool)
	var currentComponent int

	for _, n := range sorted {
		if visited[n] {
			continue
		}

		currentComponent++
		visited[n] = true
		stack := []*Node{n}

		for len(stack) != 0 {
			current := stack[len(stack) - 1]
			components[currentComponent] = append(components[currentComponent], current)
			stack = stack[:len(stack) - 1]

			for _, child := range current.out {
				if !visited[child] {
					stack = append(stack, child)
					visited[child] = true
				}
			}
		}
	}

	return components
}

func SCC(nodes []*Node) map[int][]*Node {
	// Topologically sorted based on inverse graph
	// What we want is an array of Nodes based on inverse Topo Sort
	sorted := make([]*Node, len(nodes))
	InverseTopoSort(nodes, sorted)
	return ConnectedComponents(sorted)
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
	AddEdge(b, c)
	AddEdge(c, g)
	AddEdge(g, b)
	AddEdge(c, e)
	AddEdge(e, f)
	AddEdge(f, d)
	AddEdge(d, e)
	AddEdge(e, h)
	AddEdge(h, i)
	AddEdge(i, j)
	AddEdge(j, h)

	nodes := []*Node{a, b, c, d, e, f, g, h, i, j}
	sccs := SCC(nodes)

	fmt.Printf("Got SCCs: %v\n", sccs)
}
