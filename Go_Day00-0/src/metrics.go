package main

import (
	"fmt"
	"math"
)

type metrics struct {
	mean   float64
	SD     float64
	median float64
	mode   int64
}

func newMetrics(nums []int64) (*metrics, error) {
	if len(nums) == 0 {
		return &metrics{}, nil
	}

	if err := validateBounds(nums); err != nil {
		return nil, err
	}

	var metrics = &metrics{}
	metrics.mean = mean(nums)
	metrics.median = median(nums)
	metrics.mode = mode(nums)
	metrics.SD = standardDeviation(nums, metrics.mean)

	return metrics, nil
}

func validateBounds(nums []int64) error {
	for _, num := range nums {
		if num < -100000 || num > 100000 {
			return fmt.Errorf("number %v is out of bounds of 100000 and -100000", num)
		}
	}

	return nil
}

func mean(nums []int64) float64 {
	result := sum(nums)

	return float64(result) / float64(len(nums))
}

func median(nums []int64) float64 {
	sortInt64(nums)

	if len(nums)%2 == 1 {
		return float64(nums[len(nums)/2])
	}

	leftNum := float64(nums[(len(nums)-1)/2])
	rightNum := float64(nums[(len(nums)-1)/2+1])
	return (leftNum + rightNum) / 2
}

func mode(nums []int64) int64 {
	occurrenceMap := make(map[int64]int)

	var mostFrequentNum int64
	for _, num := range nums {
		occurrenceMap[num] += 1
		if occurrenceMap[mostFrequentNum] < occurrenceMap[num] {
			mostFrequentNum = num
		} else if occurrenceMap[mostFrequentNum] == occurrenceMap[num] {
			if num < mostFrequentNum {
				mostFrequentNum = num
			}
		}
	}

	return mostFrequentNum
}

func standardDeviation(nums []int64, mean float64) float64 {
	var sum float64

	for _, num := range nums {
		sum += math.Pow(float64(num)-mean, 2)
	}

	return math.Sqrt(sum / float64(len(nums)))
}
