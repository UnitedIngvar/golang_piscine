package main

import "fmt"

func main() {
	nums, err := scanNumbers()
	if err != nil {
		fmt.Print(err)
		return
	}

	metricIds, err := scanMetricIds()
	if err != nil {
		fmt.Print(err)
		return
	}

	metrics, err := newMetrics(nums)
	if err != nil {
		fmt.Print(err)
		return
	}

	metricsPrinter, err := newMetricsPrinter(metricIds)
	if err != nil {
		fmt.Print(err)
		return
	}

	metricsPrinter.Print(metrics)
}
