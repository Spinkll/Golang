package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(MaxSum([]int{1, 2, 3, 4}, []int{1, 2, 3, 5}))
}

func MaxSum(nums1, nums2 []int) []int {

	result_first, result_second := 0, 0

	go func() {
		for i := 0; i < len(nums1); i++ {
			result_first += nums1[i]
		}
	}()
	go func() {
		for i := 0; i < len(nums2); i++ {
			result_second += nums2[i]
		}
	}()

	time.Sleep(100 * time.Millisecond)

	if result_first >= result_second {
		return nums1
	}
	return nums2
}
