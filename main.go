package main

import (
	"fmt"

	"github.com/BlueGob/Linear-regression-Go/regression"
)

func main() {
	lr := regression.NewLinearRegression("employee_salary.csv", "Salary", 0.2)
	lr.Fit()
	fmt.Println(lr.Predict(2))
	fmt.Println(lr.Mae())
	fmt.Println(lr.Mse())
	fmt.Println(lr.R2())
}
