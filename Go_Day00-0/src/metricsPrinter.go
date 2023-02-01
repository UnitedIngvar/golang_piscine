package main

import (
	"fmt"
	"strings"
)

type MetricName int32

const (
	None   MetricName = 0
	Mean   MetricName = 1
	Mode   MetricName = 2
	Median MetricName = 3
	SD     MetricName = 4
)

type metricsPrinter struct {
	metricsToPrint []MetricName
	printAll       bool
}

func newMetricsPrinter(metricIds []int64) (*metricsPrinter, error) {
	if len(metricIds) == 0 {
		return &metricsPrinter{printAll: true}, nil
	}

	if len(metricIds) > 3 {
		return nil, fmt.Errorf("you can provide a maximum of three metric names to print")
	}

	printer := &metricsPrinter{
		metricsToPrint: make([]MetricName, None, len(metricIds)),
	}

	for _, id := range metricIds {
		metricName, err := NewMetricName(id)
		if err != nil {
			return nil, err
		}

		printer.metricsToPrint = append(printer.metricsToPrint, metricName)
	}

	return printer, nil
}

func (printer *metricsPrinter) Print(metrics *metrics) {
	if printer.printAll {
		fmt.Printf(
			"Mean: %.2f,\nMode: %v\nMedian: %.2f\nSD: %.2f\n",
			metrics.mean,
			metrics.mode,
			metrics.median,
			metrics.SD)
	}

	var stringBuilder strings.Builder
	const lineLength = 15
	stringBuilder.Grow(len(printer.metricsToPrint) * lineLength)

	if printer.shouldPrint(Mean) {
		fmt.Fprintf(&stringBuilder, "Mean: %.2f\n", metrics.mean)
	}
	if printer.shouldPrint(Mode) {
		fmt.Fprintf(&stringBuilder, "Mode: %v\n", metrics.mode)
	}
	if printer.shouldPrint(Median) {
		fmt.Fprintf(&stringBuilder, "Median: %.2f\n", metrics.median)
	}
	if printer.shouldPrint(SD) {
		fmt.Fprintf(&stringBuilder, "SD: %.2f\n", metrics.SD)
	}

	fmt.Print(stringBuilder.String())
}

func (printer *metricsPrinter) shouldPrint(metricName MetricName) bool {
	for _, name := range printer.metricsToPrint {
		if metricName == name {
			return true
		}
	}

	return false
}

func NewMetricName(id int64) (MetricName, error) {
	resultName := MetricName(id)
	if resultName != Mean && resultName != Mode &&
		resultName != Median && resultName != SD {
		return None, fmt.Errorf("value %v does not represent a metric name", id)
	}

	return resultName, nil
}
