package calculation

import (
	"fmt"
)

// TakeNumbers takes input from the user
// [Q]: why are you passing "num" as a parameter? What's the purpose?
// FIXME: this function has nothing to do with the calculation. Could be moved in another pkg (something like "utils") or put as a function in the main.go file.
// BUG: handle conversion between strings and float64 values
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
