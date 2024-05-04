// FIXME: if you enter '5' as the operationChoice the program goes ahead.
// FIXME: if you do multiple operation it doesn't calculate properly
// FIXME: handle the goodness of the user's values. "d" is not a valid input here.
// FIXME: use module awareness. This means that you've to initialize a Go module.
// FIXME: move the calculation logic into another pkg.

package main

import (
	"cli/calculation"
	"fmt"
	"os"
)

func main() {
	var operationChoice int
	var number1, number2, result float64
	result = 0

	fmt.Println("\n\nWelcome to the GO CLI Calculator. Following are the options:")

	for {
		fmt.Print("\n1. Addition\n2. Subtraction\n3. Multiplication\n4. Division\n5. End\n\n")
		fmt.Print("Enter operation: ")

		if _, err := fmt.Scan(&operationChoice); err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		switch operationChoice {
		case 1, 2, 3, 4:
			number1, number2 = calculation.TakeNumbers(result)
			switch operationChoice {
			case 1:
				result = calculation.Add(number1, number2)
			case 2:
				result = calculation.Sub(number1, number2)
			case 3:
				result = calculation.Mul(number1, number2)
			case 4:
				if number2 == 0 {
					fmt.Println("Cannot divide by zero!")
					continue
				}
				result = calculation.Div(number1, number2)
			}
			fmt.Println("\n\nThe result is ", result)
		case 5:
			fmt.Println("Terminating Calculator!...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice! Please enter a number between 1 and 5.")
		}
	}
}
