package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func DisplayColumnStats(data [][]string) {
	if len(data) == 0 {
		fmt.Println("Dataset is empty")
		return
	}

	// Assuming first row is the header
	headers := data[0]
	columnCount := len(headers)

	// Print table header
	fmt.Printf("%-20s%-20s%-20s%-20s\n", "Column", "Mean", "Unique Count", "Mode")

	for col := 0; col < columnCount; col++ {
		columnData := extractColumn(data, col)
		mean := calculateMean(columnData)
		uniqueCount := len(countUnique(columnData))
		mode := calculateMode(columnData)

		// Print stats for each column
		fmt.Printf("%-20s%-20.2f%-20d%-20s\n", headers[col], mean, uniqueCount, mode)
	}
}

// extractColumn retrieves all the values from a specific column
func extractColumn(data [][]string, col int) []string {
	var column []string
	for i := 1; i < len(data); i++ { // Skip the header row
		column = append(column, strings.TrimSpace(data[i][col]))
	}
	return column
}

// calculateMean computes the mean of a numeric column
func calculateMean(column []string) float64 {
	var sum float64
	var count float64
	for _, val := range column {
		if num, err := strconv.ParseFloat(val, 64); err == nil {
			sum += num
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / count
}

// countUnique returns a map with unique values and their frequencies
func countUnique(column []string) map[string]int {
	uniqueCounts := make(map[string]int)
	for _, val := range column {
		uniqueCounts[val]++
	}
	return uniqueCounts
}

// calculateMode returns the mode of a column (the most frequent value)
func calculateMode(column []string) string {
	uniqueCounts := countUnique(column)

	// Find the mode (value with the highest frequency)
	var mode string
	maxCount := 0
	for val, count := range uniqueCounts {
		if count > maxCount {
			maxCount = count
			mode = val
		}
	}

	return mode
}
