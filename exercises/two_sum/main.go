package main

import "fmt"

// Given an array of integers, return the indices of two numbers that add up to a number we're looking for
// You can assume that the each input has exactly one answer

func bruteForce(vals []int, x int) [2]int {
	var found [2]int

	Outer:
	for i, v := range vals {
		for j, v2 := range vals[(i + 1):] {
			if (v + v2) == x {
				jIndex := i + j + 1
				fmt.Printf("v: %d, v2: %d\nGot: %d, %d\n", v, v2, i, jIndex)
				found = [2]int{i, jIndex}
				break Outer
			}
		}
	}

	return found
}

func hashMapImpl(vals []int, x int) [2]int {
	m := make(map[int]int) // { value: index }

	for i, v := range vals {
		difference := x - v

		if j, ok := m[difference]; ok {
			return [2]int{j, i}
		}

		m[v] = i
	}

	return [2]int{}
}

func main() {
 vals := []int{ 1, 2, 3, 8, 9, 10, 7, 6, 3, 2 }
 x := 12

 found := bruteForce(vals, x)
 fmt.Printf("Found: %v\n", found)

 found = hashMapImpl(vals, x)
 fmt.Printf("Hashmap impl Found: %v\n", found)
}
