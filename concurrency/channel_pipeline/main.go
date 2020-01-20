package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 1; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for num := range naturals {
			squares <- num*num
		}
		close(squares)
	}()

	for square := range squares {
		fmt.Println(square)
	}
}
