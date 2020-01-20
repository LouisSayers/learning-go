package main

import (
	"fmt"
	"sync"
)

var count int
var lock sync.Mutex

func increment() {
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Printf("INCREMENT: Count is now %d\n", count)
}

func decrement() {
	lock.Lock()
	defer lock.Unlock()
	count--
	fmt.Printf("DECREMENT: Count is now %d\n", count)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}

	wg.Wait()
	fmt.Println("DONE")
}
