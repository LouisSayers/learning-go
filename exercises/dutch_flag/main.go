package main

import "fmt"

func setVals(arr []int, reps, start, val int) {
	for i := 0; i < reps; i++ {
		arr[i + start] = val
	}
}

func sort(arr []int) {
	counts := make([]int, 3)
	for _, v := range arr {
		counts[v - 1] += 1
	}

	setVals(arr, counts[0], 0, 1)
	setVals(arr, counts[1], counts[0], 2)
	setVals(arr, counts[2], counts[0] + counts[1], 3)
}

func sort2(arr []int) {
	i := 0
	p1 := 0
	p2 := len(arr) - 1

	for i <= p2 {
		if arr[i] == 1 {
			arr[i], arr[p1] = arr[p1], arr[i]
			p1++
			i++
		} else if arr[i] == 2 {
			i++
		} else {
			arr[i], arr[p2] = arr[p2], arr[i]
			p2--
		}
	}
}

func main() {
	arr := []int{1,2,1,3,1,2,2,3,3,2,1}
	sort(arr)
	fmt.Printf("Result: %v\n", arr)

	arr2 := []int{1,2,1,3,1,2,2,3,3,2,1}
	sort2(arr2)
	fmt.Printf("Result: %v\n", arr2)
}
