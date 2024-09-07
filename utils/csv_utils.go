package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Read_csv(filename string) map[string][]float64 {
	file, err := os.Open(filename)
	if err != nil {
		panic("file not found")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var columns = []string{}
	if scanner.Scan() {
		columns_name := scanner.Text()
		columns = strings.Split(columns_name, ",")
	}
	csv := make(map[string][]float64, len(columns))

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for i := 0; i < len(line); i++ {

			num, err := strconv.ParseFloat(line[i], 64)
			if err != nil {
				panic("not a number")
			}
			csv[columns[i]] = append(csv[columns[i]], num)
		}
	}
	return csv
}
