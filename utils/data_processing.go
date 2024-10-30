package utils

import (
	"strconv"
)

func BinaryLabelling[T comparable](data [][]T) [][]int {
	var newData [][]int

	for _, row := range data {
		temp := make([]int, len(row))
		for j, val := range row {
			switch v := any(val).(type) {
			case string:
				if num, err := strconv.Atoi(v); err == nil {
					temp[j] = num
				} else if v == "Positive" || v == "Yes" || v == "Male" {
					temp[j] = 1
				} else {
					temp[j] = 0
				}
				break
			default:
				temp[j] = 0
			}
		}
		newData = append(newData, temp)
	}

	return newData
}

func ConvertToDF(data [][]int, headers []string) map[string][]int {
	df := make(map[string][]int)
	for i, val := range headers {
		length := len(data)
		tmp := make([]int, length)
		for j := 0; j < length; j++ {
			tmp[j] = data[j][i]
		}
		df[val] = tmp
	}
	return df
}
