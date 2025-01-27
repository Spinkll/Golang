package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 4; i++ {
		go func(index int) {
			results := work(index)
			for _, result := range results {
				fmt.Println(result)
			}
		}(i)
	}

	time.Sleep(1 * time.Second)
}

func work(index int) []string {
	counter := 0
	var results []string

	for i := 0; i < 6; i++ {
		counter = i
		results = append(results, fmt.Sprintf("Goroutine %d: counter = %d", index, counter))
	}

	return results
}
