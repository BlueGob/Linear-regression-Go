package regression

import (
	"lr/utils"
	"math/rand"
)

type LinearRegression struct {
	file      string
	y_column  string
	test_size float64
}
type model struct {
	slope     float64
	intercept float64
	y_mean    float64
	x_mean    float64
}

type dataSplit struct {
	x_train []float64
	x_test  []float64
	y_train []float64
	y_test  []float64
}

var data dataSplit
var m model

func NewLinearRegression(file, y_column string, test_size float64) LinearRegression {
	return LinearRegression{
		file:      file,
		y_column:  y_column,
		test_size: test_size,
	}
}

func (l LinearRegression) Fit() {
	handleErrors(&l.file, &l.y_column)
	csv := utils.Read_csv(l.file)
	_, exist := csv[l.y_column]
	if !exist {
		panic("Column not found")
	}
	y := csv[l.y_column]
	var x []float64
	delete(csv, l.y_column)
	for k := range csv {
		x = csv[k]
		break
	}
	data.x_train, data.y_train, data.x_test, data.y_test = data_split(x, y, l.test_size)
	m.x_mean = calculate_mean(&data.x_train)
	m.y_mean = calculate_mean(&data.y_train)

	m.slope = calculate_slope(&data.x_train, &data.y_train, m.x_mean, m.y_mean)
	m.intercept = m.y_mean - m.slope*m.x_mean
}

func (l LinearRegression) Predict(predict float64) float64 {
	return (m.intercept + m.slope*predict)
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

func (l LinearRegression) Mae() float64 {
	var sum float64 = 0
	for i := 0; i < len(data.y_test); i++ {
		var value float64 = (data.y_test)[i] - prediction(m.slope, m.intercept, (data.x_test)[i])
		if value < 0 {
			value = -value
		}
		sum += value
	}
	return sum / float64(len(data.y_test))
}

func (l LinearRegression) Mse() float64 {
	var sum float64 = 0
	for i := 0; i < len(data.y_test); i++ {
		var value float64 = (data.y_test)[i] - prediction(m.slope, m.intercept, (data.x_test)[i])
		sum += (value * value)
	}
	return sum / float64(len(data.y_test))
}

func (l LinearRegression) R2() float64 {
	var numerator float64 = 0
	var denominator float64 = 0
	for i := 0; i < len(data.y_test); i++ {
		var value float64 = (data.y_test)[i] - prediction(m.slope, m.intercept, (data.x_test)[i])
		numerator += (value * value)
		value = ((data.y_test)[i] - m.y_mean)
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
