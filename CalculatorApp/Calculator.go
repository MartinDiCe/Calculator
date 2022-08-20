package CalculatorApp

import "fmt"

func Sum(num1 int, num2 int) {
	var result int
	result = num1 + num2
	fmt.Printf("RESULTADO %d\n", result)
}
func Subtract(num1 int, num2 int) {
	var result int
	result = num1 - num2
	fmt.Printf("RESULTADO %d\n", result)
}
func Multiplication(num1 int, num2 int) {
	var result int
	result = num1 * num2
	fmt.Printf("RESULTADO %d\n", result)
}
func Division(num1 int, num2 int) {
	var result int

	result = num1 / num2
	fmt.Printf("RESULTADO %d\n", result)
}
