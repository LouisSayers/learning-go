package main

import "fmt"

func find(arr []int, target int) ([2]int, bool) {
	var result [2]int
	found := false

	foundIndex, ok := search(arr, target, 0, len(arr) - 1)

	if ok {
		leftMost := searchLeft(arr, target, foundIndex - 1)
		rightMost := searchRight(arr, target, foundIndex + 1)
		result, found = [2]int{leftMost, rightMost}, true
	}

	return result, found
}

func searchRight(arr []int, target int, i int) int {
	rightmost := i - 1

	for {
		if result, found := search(arr, target, i, len(arr) - 1); found {
			rightmost = result
			i = result + 1
		} else {
			break
		}
	}

	return rightmost
}

func searchLeft(arr []int, target int, i int) int {
	leftmost := i + 1

	for {
		if result, found := search(arr, target, 0, i); found {
			leftmost = result
			i = result - 1
		} else {
			break
		}
	}

	return leftmost
}

func search(arr []int, target int, i int, j int) (int, bool) {
	if j < i || i > j {
		return -1, false
	}

	midIndex := i + ((j - i) / 2.0)
	val := arr[midIndex]

	if val == target {
		return midIndex, true
	} else if target < val {
		return search(arr, target, i, midIndex - 1)
	} else {
		return search(arr, target, midIndex + 1, j)
	}
}

func main() {
	arr := []int{1,2,3,4,5,6,6,6,7,8,9,9,9,9,9,10,12} // sorted
	target := 8

	if res, ok := find(arr, target); ok {
		fmt.Printf("Found indices %v\n", res)
	} else {
		fmt.Println("No target found")
	}
}

