package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,4,5,6,6,6,6,7,8,8,8,8,8,8,9,9}
	target := 4
	// => 3, 4

	result, found := search(nums, target)
	if found {
		fmt.Println("Result of search: ", result)
	} else {
		fmt.Println("Target not found...")
	}
}

type searchDir string
const (
	SearchLeft searchDir = "Left"
	SearchRight searchDir = "Right"
)

func search(nums []int, target int) ([2]int, bool) {
	if leftmost, found := find(nums, target, 0, len(nums) - 1, SearchLeft); found {
		rightmost, _ := find(nums, target, leftmost, len(nums) - 1, SearchRight)
		return [2]int{leftmost, rightmost}, true
	} else {
		return [2]int{}, false
	}
}

func indexesFor(mid, i, j int, dir searchDir) (nextI, nextJ int) {
	switch dir {
	case SearchRight:
		nextI, nextJ = mid + 1, j
	case SearchLeft:

		nextI, nextJ = i, mid - 1
	}
	return
}

func find(nums []int, target, i, j int, dir searchDir) (int, bool) {
	if j < i || i > j {
		return -1, false
	}

	var foundIndex int
	var found bool
	midIndex := i + ((j - i) / 2)
	midVal := nums[midIndex]

	if midVal == target {
		found = true
		nextI, nextJ := indexesFor(midIndex, i, j, dir)
		if index, ok := find(nums, target, nextI, nextJ, dir); ok {
			foundIndex = index
		}	else {
			foundIndex = midIndex
		}
	} else if target < midVal  {
		foundIndex, found = find(nums, target, i, midIndex - 1, dir)
	} else { // target > midVal
		foundIndex, found = find(nums, target, midIndex + 1, j, dir)
	}

	return foundIndex, found
}

