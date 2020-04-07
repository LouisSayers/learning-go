package main

import (
	"fmt"
)

func permute(results [][]rune, arr []rune) {
	performPermute(0, results, arr, 0)
}

func performPermute(nextResultIndex int, results [][]rune, arr []rune, low int) int {
	if low >= (len(arr) - 1) {
		perm := make([]rune, len(arr))
		copy(perm, arr)
		results[nextResultIndex] = perm
		nextResultIndex += 1
		return nextResultIndex
	}

	for i := low; i < len(arr); i++ {
		arr[low], arr[i] = arr[i], arr[low]
		nextResultIndex = performPermute(nextResultIndex, results, arr, low + 1)
		arr[low], arr[i] = arr[i], arr[low]
	}

	return nextResultIndex
}

func permute2(arr []rune) [][]rune {
	results := [][]rune{}
	return performPermute2(results, arr, 0)
}

func performPermute2(results [][]rune, arr []rune, low int) [][]rune {
	if low >= (len(arr) - 1) {
		perm := make([]rune, len(arr))
		copy(perm, arr)
		results = append(results, perm)
		return results
	}

	for i := low; i < len(arr); i++ {
		arr[low], arr[i] = arr[i], arr[low]
		results = performPermute2(results, arr, low + 1)
		arr[low], arr[i] = arr[i], arr[low]
	}

	return results
}

func factorial(n int) int {
	res := 1

	for i := 2; i <= n; i += 1 {
		res = res * i
	}

	return res
}

func main() {
	arr := []rune{'a', 'b', 'c', 'd'}
	results := make([][]rune, factorial(len(arr)))
	fmt.Println("Factorial is: ", factorial(len(arr)))
	permute(results, arr)

	fmt.Printf("Results: %c\n", results)
	fmt.Printf("Results length: %d\n", len(results))

	arr2 := []rune{'d', 'e', 'f'}
	fmt.Println("Factorial is: ", factorial(len(arr2)))
	results2 := permute2(arr2)

	fmt.Printf("Results: %c\n", results2)
	fmt.Printf("Results length: %d\n", len(results2))
}
