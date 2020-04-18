package main

import "fmt"

// n courses you have to take 0 to n-1
// Some courses may have prerequisites [0,1] indicates need to do course 1 before 0
// Given the total number of courses and a list of prereqs can you finish ALL courses?
// e.g. 2, [[1,0]] => true
// e.g. 2, [[1,0], [0, 1]] => false as there is a loop

// determine if there is a loop
// method #1 loop through each node and run dfs from there => if we encounter a node that's been visited already we've found a loop
// method #2 run a topological sort algorithm, if the #nodes visited > #nodes then there is a loop
// method #3 White (unvisited), Grey (visiting), Black (visited)
// while performing DFS if we find a grey node then we've found a loop

type Edge struct {
	n1 *Node
	n2 *Node
}

func NewEdge(n, n2 *Node) *Edge {
	return &Edge{n, n2}
}

type Node struct {
	val int
	edges []*Edge
}

func (n *Node) AddEdge(n2 *Node) {
	edge := NewEdge(n, n2)
	n.edges = append(n.edges, edge)
}

func NewNode(val int) *Node {
	return &Node{val, []*Edge{}}
}

func BuildNodes(arr [][]int) map[int]*Node {
	nodes := make(map[int]*Node)

	for _, prereqs := range	arr {
		n1Index, n2Index := prereqs[0], prereqs[1]
		if nodes[n1Index] == nil {
			nodes[n1Index] = NewNode(n1Index)
		}
		if nodes[n2Index] == nil {
			nodes[n2Index] = NewNode(n2Index)
		}
		nodes[n1Index].AddEdge(nodes[n2Index])
	}

	return nodes
}

func FindLoop(n *Node, finds []bool) bool {
	if (n == nil) {
		return false
	}

	finds[n.val] = true

	for _, edge := range n.edges {
		child := edge.n2
		if finds[child.val] {
			return true
		}

		if FindLoop(child, finds) {
			return true
		}
	}

	return false
}

func method1(arr [][]int) bool {
	nodes := BuildNodes(arr)

	for _, n := range nodes {
		if loopFound := FindLoop(n, make([]bool, len(nodes))); loopFound {
			return true
		}
	}

	return false
}

func main() {
	noLoopEdges := [][]int{{1, 0}, {3, 1}, {4, 1}, {2, 0}, {5, 2}}
	loopEdges := [][]int{{1, 0}, {3, 1}, {4, 1}, {2, 0}, {5, 2}, {0, 5}}

	result := method1(noLoopEdges)
	fmt.Printf("Method 1 no loop result: %t\n", result)

	result = method1(loopEdges)
	fmt.Printf("Method 1 with Loop result: %t\n", result)
}
