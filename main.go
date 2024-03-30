package main

import "fmt"

func main() {
	// declaring variables
	var operationChoice int
	var number1, number2, result float64
	result = 0
	var flag int = 0

	fmt.Println("\n\nWelcome to the GoLang CLI Calculator. Following are the options: ")

	// infinite loop for continuous calculating
	for {
		fmt.Print("\n1. Addition\n2. Subtraction\n3. Multiplication\n4. Division\n\n")
		fmt.Print("Enter operation: ")
		fmt.Scan(&operationChoice)

		// flag is used to check whether to take inputs for the very first time
		// or taking input to continue with the previous calculation
		if flag == 0 {
			fmt.Print("\nEnter first number: ")
			fmt.Scan(&number1)
			fmt.Print("\nEnter second number: ")
			fmt.Scan(&number2)
		} else {
			number1 = result
			fmt.Print("\nEnter your number: ")
			fmt.Scan(&number2)
		}

		// performing calculation
		switch operationChoice {
		case 1:
			result = number1 + number2
		case 2:
			result = number1 - number2
		case 3:
			result = number1 * number2
		case 4:
			result = number1 / number2
		default:
			fmt.Println("Wrong input")
			break
		}

		fmt.Println("The result is ", result)

		// controling loop: whether to continue or break out of it
		fmt.Print("To continue enter 0, otherwise 1 to quit: ")
		fmt.Scan(&flag)
		if flag == 1 {
			break
		} else if flag == 0 {
			flag = 2
		} else {
			fmt.Println("Wrong input! Error!!")
			break
		}

	}

}
