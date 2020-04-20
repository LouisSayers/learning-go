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

func Reachable(n *Node, n2 *Node, visited map[*Node]bool) bool {
	if visited[n] {
		return false
	}

	visited[n] = true

	if n == n2 {
		return true
	}

	for _, child := range n.out {
    if Reachable(child, n2, visited) {
    	return true
    }
	}

	return false
}

func ComputerSCCs(nodes []*Node) map[int][]*Node {
	results := make(map[int][]*Node)
	computed := make(map[*Node]bool)
	currentCC := 0

	for i, n := range nodes {
		if computed[n] {
			continue
		}
		currentCC++
		results[currentCC] = []*Node{n}
		computed[n] = true

		for _, n2 := range nodes[(i + 1):] {
			var nToN2, n2ToN bool
			visited := make(map[*Node]bool)
		  nToN2 = Reachable(n, n2, visited)

		  if nToN2 {
			  visited = make(map[*Node]bool)
			  n2ToN = Reachable(n2, n, visited)
		  }

			if nToN2 && n2ToN {
			  results[currentCC] = append(results[currentCC], n2)
				computed[n2] = true
			}
		}
	}

	return results
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
	AddEdge(h, j)
	AddEdge(j, i)
	AddEdge(i, h)

	nodes := []*Node{a, b, c, d, e, f, g, h, i, j}
	result := ComputerSCCs(nodes)
	fmt.Printf("SCCs: %v\n", result)
}
