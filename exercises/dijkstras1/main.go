package main

import "fmt"

type Edge struct {
	weight int
	start *Node
	end *Node
}

func NewEdge(weight int, n1, n2 *Node) *Edge {
	return &Edge{weight, n1, n2}
}

type Node struct {
	val string
	out []*Edge
}

func (n *Node) String() string {
	return n.val
}

func NewNode(s string) *Node {
	return &Node{s, []*Edge{}}
}

func AddEdge(n1, n2 *Node, weight int) {
	edge := NewEdge(weight, n1, n2)
	n1.out = append(n1.out, edge)
}

func ShortestPath(a, g *Node) []*Node {
	distances := map[*Node]int{a: 0}
	visited := make(map[*Node]bool)
	prev := map[*Node]*Node{a: nil}
	current := a

	for current != nil {
		visited[current] = true

		if current == g {
			break
		}

		currentDist := distances[current]

		// loop through edges, update if distance less than current or nil
		for _, edge := range current.out {
			dist, ok := distances[edge.end]
			if !ok || (currentDist + edge.weight) < dist {
				distances[edge.end] = currentDist + edge.weight
				prev[edge.end] = current
			}
		}

		// select lowest next distance node and set that as current
		var lowest *Node
		var lowestVal int
		for n, dist := range distances {
			if _, ok := visited[n]; ok {
				continue
			}

			if lowest == nil || dist < lowestVal {
				lowest = n
				lowestVal = dist
				continue
			}
		}

		current = lowest
	}

	if current != g {
		return []*Node{}
	}

	var path []*Node
	for current != nil {
		path = append([]*Node{current}, path...)
		current = prev[current]
	}

	return path
}

func main() {
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")
	f := NewNode("f")
	g := NewNode("g")

	AddEdge(a, b, 4)
	AddEdge(a, c, 3)
	AddEdge(b, d, 3)
	AddEdge(d, f, 1)
	AddEdge(f, e, 1)
	AddEdge(c, e, 8)
	AddEdge(e, g, 3)

	path := ShortestPath(a, g)

	fmt.Printf("Shortest Path from a to g: %v\n", path)
}
