package main

import "fmt"

type Edge struct {
	n1 *Node
	n2 *Node
}

func NewEdge(n1, n2 *Node) *Edge {
	return &Edge{n1, n2}
}

type Node struct {
	val int
	edges []*Edge
}

func (n *Node) AddChild(n2 *Node) {
	edge := NewEdge(n, n2)
	n.edges = append(n.edges, edge)
}

func NewNode(val int) *Node {
	return &Node{val, []*Edge{}}
}

func BuildGraph(arr [][]int) map[int]*Node {
	nodes := make(map[int]*Node)

	for _, edge := range arr {
		n1Val := edge[1]
		n2Val := edge[0]

		// Create the nodes if they don't already exist
		if _, ok := nodes[n1Val]; !ok {
			nodes[n1Val] = NewNode(n1Val)
		}

		if _, ok := nodes[n2Val]; !ok {
			nodes[n2Val] = NewNode(n2Val)
		}

		nodes[n1Val].AddChild(nodes[n2Val])
	}

	return nodes
}

func DFS(n *Node, visited map[*Node]bool, maxIterations int) int {
	if n == nil || visited[n] {
		return 0
	}

	count := 1
	for _, edge := range n.edges {
		count += DFS(edge.n2, visited, maxIterations - 1)
		if count > maxIterations {
			break
		}
	}

	visited[n] = true

	return count
}

func hasLoops(arr [][]int) bool {
	nodeMap := BuildGraph(arr)

	var visitedCount int
	visited := make(map[*Node]bool)
	nodeCount := len(nodeMap)

	for _, node := range nodeMap {
		visitedCount += DFS(node, visited, nodeCount - visitedCount)

		if visitedCount > nodeCount {
			return true
		}
	}

	return false
}

func main() {
	noLoopEdges := [][]int{{1, 0}, {3, 1}, {4, 1}, {2, 0}, {5, 2}}
	loopEdges := [][]int{{1, 0}, {3, 1}, {4, 1}, {2, 0}, {5, 2}, {0, 5}}

	result := hasLoops(noLoopEdges)
	fmt.Printf("HasLoops no loop result: %t\n", result)

	result = hasLoops(loopEdges)
	fmt.Printf("HasLoops with Loop result: %t\n", result)

	result = hasLoops(nil)
	fmt.Printf("HasLoops for nil: %t\n", result)
}
