package main

import "fmt"

type Stack struct {
	items []int
	maxes []int
}

func (s *Stack) Max() int {
	if len(s.maxes) == 0 {
		return 0
	}

	return s.maxes[len(s.maxes) - 1]
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
	if s.Max() <= v {
		s.maxes = append(s.maxes, v)
	}
}

func (s *Stack) Pop() (int, bool) {
	lastIndex := len(s.items) - 1

	if lastIndex < 0 {
		return -1, false
	}

	popped := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	if popped == s.Max() {
		s.maxes = s.maxes[:(len(s.maxes) - 1)]
	}

	return popped, true
}

func NewStack() *Stack {
	return new(Stack)
}

func main() {
	items := []int{1,2,5,2,6,2,3,8,4,1,3,4}

	stack := NewStack()

	for _, v := range items {
		stack.Push(v)
		fmt.Printf("i: %v, max: %v\n", v, stack.Max())
	}

	fmt.Printf("Maxes: %v\n", stack.maxes)

	fmt.Println("Popping...")
	for i := 0; i < len(items); i++ {
		if item, ok := stack.Pop(); ok {
			fmt.Printf("i: %v, max: %v\n", item, stack.Max())
		}
	}
}
