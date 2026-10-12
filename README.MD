# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

## Runtime Requirements
Golang 1.24+

## Files
- ```grade_calculator.go```: Grade calculator program
- ```grade_calculator_test.go```: Grade calculator tests
- ```go.mod```: Go module information

## How it Works
There are three types of grades, each with their own weight:
- Assignments: 50%
- Exams: 35%
- Essays: 15%

The grade calculator (grade_calculator.go) finds the average grade for each category. Then, it uses the grade type weight to calculate a final grade and assign a letter to it. The letter grades are calculated as:
- A >= 90
- B >= 80
- C >=70
- D >= 60
- F < 60

Passing requires a C or better

## How to Run Unit Tests
Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```