package main

import "fmt"

func permutations(arr []int) [][]int {
	return permute([]int{}, arr)
}

func permute(head []int, tail []int) [][]int {
	results := [][]int{}

	if len(tail) == 0 {
		results = append(results, head)
	} else {
		for i, v := range tail {
			fmt.Printf("head: %v, tail: %v\n", head, tail)
			newHead := append([]int{}, head...)
			newHead = append(newHead, v)
			newTail := append([]int{}, tail[:i]...)
			newTail = append(newTail, tail[(i+1):]...)
			fmt.Printf("Newhead: %v, Newtail: %v\n", newHead, newTail)
			results = append(results, permute(newHead, newTail)...)
		}
	}

	return results
}

func main() {
	arr := []int{1,2,3}
	res := permutations(arr)

	fmt.Printf("Permutations: \n%v\n", res)
}
