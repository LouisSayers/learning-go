package main

import "fmt"

type Node struct {
	val int
	out []*Node
}

func (n *Node) AddChild(child *Node) {
	n.out = append(n.out, child)
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.val)
}

func NewNode(val int) *Node {
	return &Node{val, []*Node{}}
}

func BuildNodes(arr [][]int) []*Node {
	nodesByVal := make(map[int]*Node)

	for _, v := range arr {
		parentVal := v[1]
		childVal := v[0]

		var parent, child *Node
		var ok bool

		if parent, ok = nodesByVal[parentVal]; !ok {
			parent = NewNode(parentVal)
			nodesByVal[parentVal] = parent
		}

		if child, ok = nodesByVal[childVal]; !ok {
			child = NewNode(childVal)
			nodesByVal[childVal] = child
		}

		parent.AddChild(child)
	}

	var i int
	result := make([]*Node, len(nodesByVal))

	for _, v := range nodesByVal {
		result[i] = v
		i++
	}

	return result
}

// Return if a loop is detected
func DFS(n *Node, inProgress, complete map[*Node]bool) bool {
	if n == nil || complete[n] {
		return false
	}

	if inProgress[n] {
		return true
	}

	inProgress[n] = true
	for _, child := range n.out {
		result := DFS(child, inProgress, complete)
		if result {
			return true
		}
	}
	complete[n] = true

	return false
}

func detectLoop(arr [][]int) bool {
	nodes := BuildNodes(arr)
	inProgress := make(map[*Node]bool)
	complete := make(map[*Node]bool)

	for _, n := range nodes {
		if complete[n] {
			continue
		}

		if DFS(n, inProgress, complete) {
			return true
		}

		// reset inProgress
		inProgress = make(map[*Node]bool)
	}

	return false
}

func main() {
	noLoopArr := [][]int{{1, 0}, {2, 1}, {3, 0}}
	hasLoops := detectLoop(noLoopArr)
	fmt.Printf("No Loop Array has a loop: %t\n", hasLoops)

	loopArr := [][]int{{1, 0}, {2, 1}, {3, 0}, {4, 3}, {5, 4}, {3, 5}}
	hasLoops = detectLoop(loopArr)
	fmt.Printf("Loop Array has a loop: %t\n", hasLoops)
}
