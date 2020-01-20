package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type balance struct {
	amount   float64
	currency string
	mu       sync.Mutex
}

func (b *balance) Add(i float64) {
	b.mu.Lock()
	amount := b.amount
	time.Sleep(1)
	newAmount := amount + i
	b.amount = newAmount
	b.mu.Unlock()
}

func (b *balance) Display() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return strconv.FormatFloat(b.amount, 'f', 2, 64) + " " + b.currency
}

func run(wg *sync.WaitGroup) float64 {
	var localWG sync.WaitGroup

	b := &balance{amount: 50, currency: "AUD"}

	wg.Add(2)
	localWG.Add(2)

	go func() {
		defer wg.Done()
		defer localWG.Done()
		b.Add(50.0)
		b.Add(50.0)
	}()

	go func() {
		defer wg.Done()
		defer localWG.Done()
		b.Add(100.0)
		b.Add(100.0)
	}()

	localWG.Wait()
	return b.amount
}

func main() {
	var wg sync.WaitGroup
	results := make(map[float64]struct{})

	for i := 0; i < 1000; i++ {
		balance := run(&wg)
		results[balance] = struct{}{}
	}

	wg.Wait()

	fmt.Println("Got results:")
	for k := range results {
		fmt.Printf(strconv.FormatFloat(k, 'f', 2, 64))
	}
	fmt.Println("")
}
