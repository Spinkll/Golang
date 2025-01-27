package main

import (
	"fmt"
)

func main() {
	fmt.Println(MaxSum([]int{1, 2, 3, 4}, []int{1, 2, 3}))
}

func MaxSum(nums1, nums2 []int) []int {

	result_first, result_second := 0, 0

	for i := 0; i < len(nums1); i++ {
		result_first += nums1[i]
	}

	for i := 0; i < len(nums2); i++ {
		result_second += nums2[i]
	}

	if result_first >= result_second {
		return nums1
	}
	return nums2
}
