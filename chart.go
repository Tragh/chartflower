package main

import "fmt"

func charts() [1]string {
	charts := [1]string{"bar"}
	return charts
}

func makeChart(data [][]string) string {
	var chart string

	switch selectChart() {
	case "bar":
		chart = barChart(data)
	default:
		fmt.Println("No case for", chart)
	}

	return chart
}

func barChart(data [][]string) string {

	return "bar chart here"
}
