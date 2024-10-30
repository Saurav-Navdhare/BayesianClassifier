package utils

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"os"
)

func convertToRow[T comparable](row []T) table.Row {
	newRow := make(table.Row, len(row))
	for i, v := range row {
		switch val := any(v).(type) {
		case string:
			newRow[i] = val
		default:
			newRow[i] = fmt.Sprintf("%v", val)
		}
	}
	return newRow
}

func DrawTableFromMatrix[T comparable](data [][]T, headers []string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Set the table headers
	t.AppendHeader(convertToRow(headers))

	for _, row := range data {
		t.AppendRow(convertToRow(row))
	}

	t.Render()
}

func DrawTable(df map[string][]int) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Set the table headers

	var tmp []string
	for key := range df {
		tmp = append(tmp, key)
	}
	t.AppendHeader(convertToRow(tmp))

	for _, v := range df[tmp[0]] {
		var tmp []interface{}
		for key := range df {
			tmp = append(tmp, df[key][v])
		}
		t.AppendRow(convertToRow(tmp))
	}

	t.Render()
}
