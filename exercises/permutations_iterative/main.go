package main

import "fmt"

func permute(arr []int) [][]int {
	seed := [][]int{{}, arr}
	stack := [][][]int{seed}
	results := [][]int{}

	for len(stack) != 0 {
		lastItemIdx := len(stack) - 1
		item := stack[lastItemIdx]
		stack = stack[:lastItemIdx]
		head, tail := item[0], item[1]

		for i, v := range tail {
			newHead := make([]int, len(head) + 1)
			copy(newHead, head)
			newHead[len(newHead) - 1] = v

			newTail := make([]int, len(tail) - 1)
			copy(newTail, tail[:i])
			copy(newTail[i:], tail[(i+1):])

			if len(newTail) == 0 {
				results = append(results, newHead)
			} else {
				newItem := [][]int{newHead, newTail}
				stack = append(stack, newItem)
			}
		}
	}

	return results
}

func main() {
	arr := []int{1,2,3, 4}
	permutations := permute(arr)

	fmt.Printf("Results: %v\n", permutations)
}
