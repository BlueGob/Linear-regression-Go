# Simple Linear Regression with Go
This project provides an implementation of linear regression with two variables in Go.
It includes various evaluation metrics such as R² (R-squared), Mean Squared Error (MSE),and Mean Absolute Error (MAE).
## Features
-  Load data from a CSV file
- Fit a linear regression model
- Calculate MAE, MSE, and R²
- Make predictions based on the model
## Getting Started

### Usage
1. ```bash
   go get github.com/BlueGob/Linear-regression-Go
   ``` 
2. Create a Go file (`main.go`) with the following content:

    ```go
    package main

    import (
        "fmt"
        "github.com/BlueGob/Linear-regression-Go/regression"
    )

    func main() {
        lr := regression.NewLinearRegression("employee_salary.csv", "Salary", 0.2)
        lr.Fit()
        fmt.Println("Prediction for input 2:", lr.Predict(2))
        fmt.Println("Mean Absolute Error (MAE):", lr.Mae())
        fmt.Println("Mean Squared Error (MSE):", lr.Mse())
        fmt.Println("R-squared (R²):", lr.R2())
    }
    ```

2. Run the program:

    ```bash
    go run main.go
    ```
### Dataset
The dataset used for this project can be found on [Kaggle](https://www.kaggle.com/datasets/hassanmustafa01/employee-salary-dataset)
