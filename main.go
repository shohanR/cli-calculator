package main

import (
	"fmt"
	"os"
)

// function to take input
func TakeNumbers(num float64) (float64, float64) {
	var num1, num2 float64
	if num == 0 {
		fmt.Print("\nEnter first number: ")
		fmt.Scan(&num1)
		fmt.Print("\nEnter second number: ")
		fmt.Scan(&num2)
	} else {
		num1 = num
		fmt.Print("Enter your number: ")
		fmt.Scan(&num2)
	}

	return num1, num2
}

// function to perform Addition
func Add(num1, num2 float64) float64 {
	return num1 + num2
}

// function to perform Subtraction
func Sub(num1, num2 float64) float64 {
	return num1 - num2
}

// function to perform Multiplication
func Mul(num1, num2 float64) float64 {
	return num1 * num2
}

// function to perform Division
func Div(num1, num2 float64) float64 {
	return num1 / num2
}

func main() {
	var operationChoice int
	var err error // error type variable initialization
	var number1, number2, result float64
	result = 0

	fmt.Println("\n\nWelcome to the GO CLI Calculator. Following are the options:")

	// infinite loop for continuous calculations
	for {
		fmt.Print("\n1. Addition\n2. Subtraction\n3. Multiplication\n4. Division\n5.End\n\n")
		fmt.Print("Enter operation: ")
		_, err = fmt.Scan(&operationChoice) //error status handling

		// error handling
		if err != nil {
			fmt.Println("Error reading input:", err)
			// os.Exit(0)
			break
		}

		if operationChoice < 1 || operationChoice > 5 {
			fmt.Print("Wrong Choice! If you wish to continue, enter 1, otherwise enter 0 to end: ")
			_, err = fmt.Scan(&operationChoice)
			if err != nil {
				fmt.Println("Error reading input:", err)
				// os.Exit(0)
				break
			}

			if operationChoice == 1 {
				continue
			} else if operationChoice == 0 {
				fmt.Println("Terminating Calculator! ...")
				// os.Exit(0)
				break
			} else {
				fmt.Println("Wrong choice! Terminating Program!")
				// os.Exit(0)
				break
			}
		} else if operationChoice == 5 {
			fmt.Println("Terminating Calculator!...")
			// os.Exit(0)
			break
		} else {
			number1, number2 = TakeNumbers(result)

			switch operationChoice {
			case 1:
				result = Add(number1, number2)
			case 2:
				result = Sub(number1, number2)
			case 3:
				result = Mul(number1, number2)
			case 4:
				result = Div(number1, number2)
			}
		}
	}

	fmt.Printf("The result is %0.2f", result)
	os.Exit(0)
}
