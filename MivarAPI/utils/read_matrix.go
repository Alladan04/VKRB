package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadMatrixFromFile(filename string) ([][]uint8, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	// Парсинг матрицы (примерная реализация)
	var matrix [][]uint8
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var row []uint8
		for _, char := range strings.Split(line, " ") {
			num, err := strconv.Atoi(char)
			if err != nil {
				return nil, fmt.Errorf("invalid matrix format: %w", err)
			}
			row = append(row, uint8(num))
		}
		matrix = append(matrix, row)
	}

	return matrix, nil
}
