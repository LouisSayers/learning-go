package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("./data/example.txt")

	if err != nil {
		fmt.Printf("Got error: '%v'\n", err.Error())
	} else {
		strContents := string(contents[:])
		parts := strings.Split(strContents, "\n")
		re := regexp.MustCompile(`\S+ \S+ (\d+)`)

		entryCount := len(parts)
		if parts[len(parts)-1] == "" {
			entryCount--
		}

		nums := make([]int, entryCount)

		for i, v := range parts {
			matches := re.FindStringSubmatch(v)
			if len(matches) > 1 {
				num, err := strconv.Atoi(matches[1])

				if err != nil {
					panic("Tried to convert a non number...")
				}

				nums[i] = num
			}
		}

		sort.Ints(nums)
		sum := 0

		for _, v := range nums {
			sum += v
		}

		fmt.Println(nums)
		fmt.Printf("Mean: %v\n", float32(sum)/float32(len(nums)))

		if len(nums)%2 == 0 {
			a := nums[len(nums)/2-1]
			b := nums[len(nums)/2]
			median := (float32(a) + float32(b)) / 2

			fmt.Printf("Median: %v\n", median)
		} else {
			fmt.Printf("Median: %v\n", nums[len(nums)/2])
		}
	}
}
