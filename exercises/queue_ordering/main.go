package main

import (
	"fmt"
	"sort"
)

type Person struct {
	h int // height
	k int // number of people in front >= height
}

func (p *Person) String() string {
	return fmt.Sprintf("h: %d, k: %d\n", p.h, p.k)
}

type People []*Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	if p[i].h == p[j].h {
		return p[i].k < p[j].k
	}

	return p[i].h >= p[j].h // Desc order
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func peformSort1(input People) {
	sort.Sort(input)

	result := make([]*Person, len(input))

	for i := 0; i < len(input); i++ {
		current := input[i]
		for j := current.k; j < (i + 1); j++ {
			temp := result[j]
			result[j] = current
			current = temp
		}
	}

	fmt.Printf("People: %v\n", result)
}

func peformSort2(input People) {
	sort.Sort(input)

	for i := 0; i < len(input); i++ {
		current := input[i]

		for j := current.k; j <= i; j++ {
			temp := input[j]
			input[j] = current
			current = temp
		}
	}

	fmt.Printf("People: %v\n", input)
}

func main() {
	input := People{{7,0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	peformSort1(input)
	peformSort2(input)
}
