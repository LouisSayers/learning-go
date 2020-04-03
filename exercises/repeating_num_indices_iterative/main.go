package main

import "fmt"

func main() {
	nums := []int{1,2,3,3,3,4,5,6,6,6,6,7,8,8,9}
	target := 6
	// expecting [7, 10]

	if indices, ok := findIndices(nums, target); ok {
		fmt.Printf("Indices: %v\n", indices)
	} else {
		fmt.Println("Target not found...")
	}
}

type findStrategy string
const (
	Any findStrategy ="ANY"
	Leftmost findStrategy = "LEFT"
	Rightmost findStrategy = "RIGHT"
)

func findIndices(nums []int, target int) ([2]int, bool) {
	foundIndex, i, j, found := find(nums, target, 0, len(nums) - 1, -1, Any)

	if found {
		leftmost, _, _, _ := find(nums, target, i, foundIndex - 1, foundIndex, Leftmost)
		rightmost, _, _, _ := find(nums, target, foundIndex + 1, j, foundIndex, Rightmost)
		return [2]int{leftmost, rightmost}, true
	} else {
		return [2]int{}, false
	}
}

func find(nums []int, target, i, j, foundIndex int, strategy findStrategy) (int, int, int, bool) {
	found := false

	OuterLoop:
	for {
		if j < i || i > j {
			break
		}
		midPoint := i + (j - i) / 2
		midVal := nums[midPoint]
		if midVal == target {
			found = true
			switch strategy {
			case Any:
				foundIndex = midPoint
				break OuterLoop
			case Leftmost:
				foundIndex = midPoint
				j = midPoint - 1
			case Rightmost:
				foundIndex = midPoint
				i = midPoint + 1
			}
		} else if target < midVal {
			j = midPoint - 1
		} else {
			i = midPoint + 1
		}
	}

	return foundIndex, i, j, found
}
