package main

import (
	"fmt"
)

func main() {
	numsCh := make(chan []int)
	sumCh := make(chan int)

	go SumWorker(numsCh, sumCh)
	numsCh <- []int{10, 10, 10}
	res := <-sumCh
	fmt.Println(res)
}

func SumWorker(numsCh chan []int, sumCh chan int) {
	for nums := range numsCh {
		result := 0
		for i := 0; i < len(nums); i++ {
			result += nums[i]
		}
		sumCh <- result
	}
}
