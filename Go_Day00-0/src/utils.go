package main

import "sort"

func sortInt64(nums []int64) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}

func sum(nums []int64) int64 {
	var result int64
	for _, i := range nums {
		result += i
	}
	return result
}
