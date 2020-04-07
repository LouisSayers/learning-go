package main

import "fmt"

type StackItem struct {
	low int
	i int
	undoA int
	undoB int
}

func permute(arr []rune) {
	stack := []StackItem{ {0, 0, 0, 0} }
	for len(stack) != 0 {
		lastStackIndex := len(stack) - 1
		current := stack[lastStackIndex]
		stack = stack[:lastStackIndex]

		lastArrIndex := len(arr) - 1
		if current.i == lastArrIndex && current.low == lastArrIndex { // complete permutation
			fmt.Printf("%c\n", arr)
			arr[current.undoA], arr[current.undoB] = arr[current.undoB], arr[current.undoA] // undo
		} else if current.i > lastArrIndex { // finished this level
			arr[current.undoA], arr[current.undoB] = arr[current.undoB], arr[current.undoA] // undo
		} else { // not finished
			arr[current.low], arr[current.i] = arr[current.i], arr[current.low] // swap values
			child := StackItem{current.low + 1, current.low + 1, current.low, current.i}
			current.i += 1
			stack = append(stack, current)
			stack = append(stack, child)
		}
	}
}

func main() {
	arr := []rune{'l', 'o', 'u', 'i', 's'}
	permute(arr)
}
