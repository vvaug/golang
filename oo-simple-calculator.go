package main

import (
	"fmt"
	"strconv"
)

type Calculator struct {
	firstNumber  int64
	secondNumber int64
	operator     string
}

func main() {
	introduction()
	var operation string
	fmt.Scan(&operation)
	validate(operation)
	calculator := createCalculator(operation)
	result := calculate(calculator)
	fmt.Println(result)
}

func introduction() {
	fmt.Println("Digit a number, a operation and another number.")
	fmt.Println("Example: 2*2")
}

func validate(operation string) {
	operator := string(operation[1])
	if operator != "+" && operator != "-" && operator != "*" && operation != "/" {
		panic("Invalid operator:" + operator)
	}
}

func createCalculator(operation string) *Calculator {
	firstNumber, _ := strconv.ParseInt(string(operation[0]), 10, 64)
	secondNumber, _ := strconv.ParseInt(string(operation[2]), 10, 64)
	calculator := Calculator{firstNumber, secondNumber, string(operation[1])}
	return &calculator
}

func calculate(calculator *Calculator) int64 {
	if calculator.operator == "*" {
		return calculator.firstNumber * calculator.secondNumber
	} else if calculator.operator == "/" {
		return calculator.firstNumber / calculator.secondNumber
	} else if calculator.operator == "+" {
		return calculator.firstNumber + calculator.secondNumber
	} else if calculator.operator == "-" {
		return calculator.firstNumber - calculator.secondNumber
	}
	panic("an error occurred while trying to do operation")
}
