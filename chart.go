package main

import (
	"fmt"
	"strconv"
	"strings"
)

func charts() [3]string {
	charts := [3]string{"bar", "pie", "radar"}
	return charts
}

func makeChart(data [][]string) string {
	var chart string

	switch selectChart() {
	case "bar":
		chart = barChart(data)
	case "pie":
		chart = pieChart(data)
	case "radar":
		chart = radarChart(data)
	default:
		fmt.Println("No case for", chart)
	}

	return chart
}

func radarChart(data [][]string) string {
	labelsIndex, labelsString := chooseLabelsColumn(data)
	valuesLabel, valuesString := chooseValuesColumn(labelsIndex, data)
	return radarChartTemplate(labelsString, valuesString, valuesLabel)
}

func pieChart(data [][]string) string {
	labelsIndex, labelsString := chooseLabelsColumn(data)
	valuesLabel, valuesString := chooseValuesColumn(labelsIndex, data)
	return pieChartTemplate(labelsString, valuesString, valuesLabel)
}

func barChart(data [][]string) string {
	labelsIndex, labelsString := chooseLabelsColumn(data)
	valuesLabel, valuesString := chooseValuesColumn(labelsIndex, data)
	return barChartTemplate(labelsString, valuesString, valuesLabel)
}

func chooseValuesColumn(labelsIndex int, data [][]string) (string, string) {
	columnNames := getColumnNames(data)
	firstRow := getFirstRow(data)
	var valuesIndex int
	var values []string

	fmt.Println()
	for i, column := range columnNames {
		if i != labelsIndex {
			fmt.Println(strconv.Itoa(i) + ". " + column + " eg. " + firstRow[i])
		}
	}

	choice := getChoice("Choose datasets")
	choices := strings.Fields(choice)
	fmt.Println(choices)
	if len(choices) == 1 {
		for i, column := range columnNames {
			number := strconv.Itoa(i)
			if choice == column {
				valuesIndex = i
				values = getColumnData(i, data)
			} else if choice == number {
				valuesIndex = i
				values = getColumnData(i, data)
			}
		}
	} else {
		// multiple datasets chosen
	}
	valuesLabel := columnNames[valuesIndex]
	valuesString := strings.Join(values, ",")
	return valuesLabel, valuesString
}

func chooseLabelsColumn(data [][]string) (int, string) {
	columnNames := getColumnNames(data)
	firstRow := getFirstRow(data)
	var labelsIndex int
	var labels []string

	fmt.Println()
	for i, column := range columnNames {
		fmt.Println(strconv.Itoa(i) + ". " + column + " eg. " + firstRow[i])
	}

	choice := getChoice("Choose label column")

	for i, column := range columnNames {
		number := strconv.Itoa(i)
		if choice == column {
			labelsIndex = i
			labels = getColumnData(i, data)
		} else if choice == number {
			labelsIndex = i
			labels = getColumnData(i, data)
		}
	}

	labelsString := strings.Join(labels, ",")
	return labelsIndex, labelsString
}

func getColumnData(column int, data [][]string) []string {
	var columnData []string
	for i, data := range data {
		if i != 0 {
			columnData = append(columnData, data[column])
		}
	}
	return columnData
}

func getFirstRow(data [][]string) []string {
	var columns []string
	for i, row := range data {
		if i == 1 {
			columns = row
		}
	}
	return columns
}

func getColumnNames(data [][]string) []string {
	var columns []string
	for i, row := range data {
		if i == 0 {
			columns = row
		}
	}
	return columns
}

func generateColors(number int) {

}
