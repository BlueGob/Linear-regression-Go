# Simple Linear Regression with Go
This project provides an implementation of linear regression with two variables in Go.
It includes various evaluation metrics such as R² (R-squared), Mean Squared Error (MSE),and Mean Absolute Error (MAE).
### Features
- **Linear Regression:** Predicts a target variable based on one feature.
- **Evaluation Metrics:** Computes R², MSE, and MAE to evaluate the model.
- **Read Directly from CSV**: Load data directly from a CSV file for regression analysis.
- **Command-Line Interface:** Allows easy configuration and execution of the regression analysis
### Installation
```bash
git clone https://github.com/BlueGob/Linear-regression-Go
```
```bash
cd Linear-regression-Go
```
- Ensure you have Go installed.

### Usage
```go
go run regression.go -file <csv-file> -y <y-column> -test_size <test-size> -predict <predict-value> -metrics <metrics>
```
#### Arguments
- `-file <csv-file>`: Path to the CSV file containing the data. This argument is mandatory.
- `-y <y-column>`: The name of the column you want to perform regression on and make predictions for. This argument is mandatory.
- `-test_size <test-size>`: Proportion of the data to be used for testing (between 0 and 1). Default is 0.2.
- `-predict <predict-value>`: The value for which you want to make a prediction using the trained model.
- `-metrics <metrics>`: Evaluation metrics to compute. Options include r2 for R-squared, mae for Mean Absolute Error, and mse for Mean Squared Error.
  Multiple metrics can be specified, separated by commas.

  #### Example
  To execute the program with the following parameters:
- `CSV File`: employee_salary.csv
- `Target Column`: Salary
- `Test Size`: 0.2
- `Prediction Value`: 5
- `Metrics`: r2, MAE, and MSE
  Use the command:
```go
go run regression.go -file employee_salary.csv -y Salary -test_size 0.2 -predict 5 -metrics=r2,mae,mse
```
### Notes
- Ensure that the CSV file and column names are correctly specified.
- The `-test_size` parameter must be between 0 and 1.
- For the `-metrics` parameter, you can specify one or more metrics, separated by commas. If left empty, all available metrics will be computed.
### Dataset
The dataset used for this project can be found on [Kaggle](https://www.kaggle.com/datasets/hassanmustafa01/employee-salary-dataset)
