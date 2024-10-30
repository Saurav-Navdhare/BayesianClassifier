package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

// LoadData loads the CSV file and returns it as a 2D slice of strings
func LoadData(filePath string) ([][]string, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read csv file: %v", err)
	}

	return records, nil
}

func Head(df map[string][]int, n int) {
	if n == 0 {
		n = 5
	}
	data := make(map[string][]int)
	for key := range df {
		n = min(n, len(df[key]))
		data[key] = df[key][:n]
	}
	DrawTable(data)
}
