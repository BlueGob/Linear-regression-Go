package main

import (
	"flag"
	"fmt"
	"math/rand"
	"root/utils"
	"strings"
)

func main() {
	file := flag.String("file", "", "Path to the CSV file")
	y_column := flag.String("y", "", "Name of the target column for regression")
	test_size := flag.Float64("test_size", 0.2, "Proportion of data to be used for testing (range: 0 to 1, default: 0.2)")
	predict := flag.Float64("predict", -1, "X value for which you want to make a prediction")
	metrics := flag.String("metrics", "", "Comma-separated evaluation metrics (e.g., 'mea,mse,r2'). Leave empty to select them all")
	flag.Parse()

	handleErrors(file, y_column)

	csv := utils.Read_csv(*file)
	_, exist := csv[*y_column]
	if !exist {
		panic("Column not found")
	}
	y := csv[*y_column]
	var x []float64
	delete(csv, *y_column)
	for k := range csv {
		x = csv[k]
		break
	}
	x_train, y_train, x_test, y_test := data_split(x, y, *test_size)
	mean_x := calculate_mean(&x_train)
	mean_y := calculate_mean(&y_train)
	var slope float64 = calculate_slope(&x_train, &y_train, mean_x, mean_y)
	var intercept = mean_y - slope*mean_x
	if *predict > 0 {
		fmt.Printf("The prediction for Value %f is : %f \n", *predict, prediction(slope, intercept, *predict))
	}
	display(metrics, &x_test, &y_test, slope, intercept, mean_y)
}

func calculate_mean(arr *[]float64) float64 {
	var sum float64 = 0
	for _, v := range *arr {
		sum += v
	}
	return sum / float64(len(*arr))
}

func calculate_slope(x *[]float64, y *[]float64, mean_x float64, mean_y float64) float64 {
	var numerator float64 = 0
	var denominator float64 = 0
	for i := 0; i < len(*x); i++ {
		numerator += ((float64((*x)[i]) - mean_x) * (float64((*y)[i]) - mean_y))
		denominator += ((float64((*x)[i]) - mean_x) * (float64((*x)[i]) - mean_x))
	}
	return numerator / denominator
}
func prediction(slope float64, intercept float64, x float64) float64 {
	return (intercept + slope*float64(x))
}
func data_split(x []float64, y []float64, test_size float64) ([]float64, []float64, []float64, []float64) {
	rand.Shuffle(len(x), func(i, j int) {
		x[i], x[j] = x[j], x[i]
		y[i], y[j] = y[j], y[i]
	})
	var n int = int(float64(len(x)) * test_size)
	x_test := x[0:n]
	x_train := x[n:]
	y_test := y[0:n]
	y_train := y[n:]
	return x_train, y_train, x_test, y_test
}

func mae(y_test *[]float64, x_test *[]float64, slope float64, intercept float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(*y_test); i++ {
		var value float64 = (*y_test)[i] - prediction(slope, intercept, (*x_test)[i])
		if value < 0 {
			value = -value
		}
		sum += value
	}
	return sum / float64(len(*y_test))
}

func mse(y_test *[]float64, x_test *[]float64, slope float64, intercept float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(*y_test); i++ {
		var value float64 = (*y_test)[i] - prediction(slope, intercept, (*x_test)[i])
		sum += (value * value)
	}
	return sum / float64(len(*y_test))
}

func r2(y_test *[]float64, x_test *[]float64, slope float64, intercept float64, y_mean float64) float64 {
	var numerator float64 = 0
	var denominator float64 = 0
	for i := 0; i < len(*y_test); i++ {
		var value float64 = (*y_test)[i] - prediction(slope, intercept, (*x_test)[i])
		numerator += (value * value)
		value = ((*y_test)[i] - y_mean)
		denominator += (value * value)

	}
	return 1 - numerator/denominator
}
func handleErrors(file *string, y_column *string) {
	if *file == "" {
		panic("csv file not provided")
	}
	if *y_column == "" {
		panic("y_column not provided")
	}
}
func display(metrics *string, x_test *[]float64, y_test *[]float64, slope float64, intercept float64, mean_y float64) {
	if *metrics == "" {
		fmt.Println("mae = ", mae(y_test, x_test, slope, intercept))
		fmt.Println("mse = ", mse(y_test, x_test, slope, intercept))
		fmt.Println("r2 = ", r2(y_test, x_test, slope, intercept, mean_y))
	} else {
		metric_selection := strings.Split(*metrics, ",")
		for _, v := range metric_selection {
			switch v {
			case "mae":
				fmt.Println("mae = ", mae(y_test, x_test, slope, intercept))
			case "r2":
				fmt.Println("r2 = ", r2(y_test, x_test, slope, intercept, mean_y))
			case "mse":
				fmt.Println("mse = ", mse(y_test, x_test, slope, intercept))
			}
		}
	}
}
