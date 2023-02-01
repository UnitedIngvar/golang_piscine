package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func scanNumbers() ([]int64, error) {
	fmt.Print("Please, provide a set of numbers, separated by new lines " +
		"(press ctrl + d (ctrl + z on windows) to stop):\n")
	nums, err := scanIntArray()
	if err != nil {
		return nil, err
	}

	return nums, nil
}

func scanMetricIds() ([]int64, error) {
	fmt.Print("Please, type the numbers, separated by newlines, " +
		"for metrics you need to see (press ctrl + d (ctrl + z on windows) to stop). \n" +
		"Press enter to print all metrics: \n" +
		"1 - Mean\n" +
		"2 - Mode\n" +
		"3 - Median\n" +
		"4 - Standard Deviation\n")

	metricIds, err := scanIntArray()
	if err != nil {
		return nil, err
	}

	return metricIds, nil
}

func scanIntArray() ([]int64, error) {
	const defaultArrayLen = 100
	const defaultValue = 0
	result := make([]int64, defaultValue, defaultArrayLen)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return nil, err
		}

		result = append(result, number)
	}

	return result, nil
}
