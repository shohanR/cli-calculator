package calculation

import (
	"fmt"
)

// TakeNumbers takes input from the user
func TakeNumbers(num float64) (float64, float64) {
	var num1, num2 float64
	if num == 0 {
		fmt.Print("\nEnter first number: ")
		fmt.Scan(&num1)
		fmt.Print("\nEnter second number: ")
		fmt.Scan(&num2)
	} else {
		num1 = num
		fmt.Print("Enter second number: ")
		fmt.Scan(&num2)
	}

	return num1, num2
}

// Add performs addition
func Add(num1, num2 float64) float64 {
	return num1 + num2
}

// Sub performs subtraction
func Sub(num1, num2 float64) float64 {
	return num1 - num2
}

// Mul performs multiplication
func Mul(num1, num2 float64) float64 {
	return num1 * num2
}

// Div performs division
func Div(num1, num2 float64) float64 {
	if num2 == 0 {
		return 0
	}
	return num1 / num2
}
