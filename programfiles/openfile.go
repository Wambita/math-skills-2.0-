package mathskills

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func OpenFile(filePath string) ([]float64, error) {
	inputFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file")
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var data []float64
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("could not convert string to float: %v", err)
		}
		if num >= float64(math.MaxInt64) {
			return nil, fmt.Errorf("value out of int64 range")
		}
		if num < 0 {
			return nil, fmt.Errorf("population contains a negative value")
		}

		data = append(data, num)
		if len(data) < 2 {
			return nil, fmt.Errorf("population has one value, add more values")
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
