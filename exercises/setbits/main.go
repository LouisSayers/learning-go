package main

import "fmt"

func shiftCount(i int) int {
	count := 0
	for ; i != 0; i >>= 1 {
		count += i & 1
	}
	return count
}

func shiftCountRecursive(i int) int {
	if i == 0 {
		return 0
	}

	return i & 1 + shiftCountRecursive(i>>1)
}

func kernighansCount(i int) int {
	count := 0
	for i != 0 {
		i = i & (i - 1)
		count++
	}
	return count
}

func kernighansCountRecursive(i int) int {
	if i == 0 {
		return 0
	}

	return 1 + kernighansCountRecursive(i & (i - 1))
}

var calcs [256]uint8

func init() {
	for i := 0; i < 256; i++ {
		calcs[i] = calcs[i/2] + uint8(i & 1)
	}
}

func bitset(orig uint64) int {
	var sum int
	sum = int(calcs[uint8(orig >> 0 * 8)]) +
		int(calcs[uint8(orig >> (1 * 8))]) +
		int(calcs[uint8(orig >> (2 * 8))]) +
		
		int(calcs[uint8(orig >> (3 * 8))]) +
		int(calcs[uint8(orig >> (4 * 8))]) +
		int(calcs[uint8(orig >> (5 * 8))]) +
		int(calcs[uint8(orig >> (6 * 8))]) +
		int(calcs[uint8(orig >> (7 * 8))])

	return sum
}

func main() {
	res := shiftCount(7)
	fmt.Printf("Bit count: %d\n", res)

	res = shiftCountRecursive(9)
	fmt.Printf("Bit count: %d\n", res)

	res = kernighansCount(7)
	fmt.Printf("Bit count: %d\n", res)

	res = kernighansCountRecursive(9)
	fmt.Printf("Bit count: %d\n", res)

	res = bitset(15)
	fmt.Printf("Bit count: %d\n", res)
}
