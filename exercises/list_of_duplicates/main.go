package main

import "fmt"

func hashMapSolution(arr []int) int {
	count := make(map[int]int)
	for _, v := range arr {
		count[v] += 1
	}

	var res int

	for k, v := range count {
		if v == 1 {
			res = k
			break
		}
	}

	return res
}

func xorSolution(arr []int) int {
	var result int

	for _, v := range arr {
		result ^= v
	}

	return result
}

func main() {
	arr := []int{1, 2, 6, 3, 3, 2, 1}

	result := hashMapSolution(arr)
	fmt.Printf("Result: %d\n", result)

	arr = []int{1, 2, 6, 3, 6, 3, -10, 2, 1}
	result = xorSolution(arr)
	fmt.Printf("Result: %d\n", result)
}
