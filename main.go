// FIXME: handle the goodness of the user's values. "d" is not a valid input here.
// TODO: let's do only one operation at time. Every time you want to perform an operation asks for all the data you need. NO OPERATION ON THE PREVIOUS OPERATION'S RESULT.
package main

import (
	"fmt"
	"os"
	"strconv"

	"cli-calculator/calculation"
	"cli-calculator/utils"
)

// TODO: try to remove logic from the main() since it's too much bloated with logic, checks, and so on.
func main() {
	var result float64 = 0
	// TODO: "checker" identifier is not meaningful. What's its purpose?
	checker := 0

	fmt.Println("\nWelcome to the GO CLI Calculator. Following are the options:")

	for {
		var operationChoice string
		var number1, number2 float64
		flag := 0

		if checker > 0 {
			fmt.Print("Would you like to continue with the current calculation (Enter '1'), or start a new one (Enter '0')? ")

			var continueChoice string
			fmt.Scan(&continueChoice)

			temp, err := strconv.Atoi(continueChoice)
			if err != nil || (temp < 0 || temp > 1) {
				fmt.Println("Invalid input! Please enter 0 or 1.")
				continue
			}

			if temp == 0 {
				result = 0
			}
		}

		fmt.Print("\n1. Addition\n2. Subtraction\n3. Multiplication\n4. Division\n5. End\n")
		fmt.Print("Enter operation: ")
		fmt.Scan(&operationChoice)

		opChoiceInt, err := strconv.Atoi(operationChoice)
		if err != nil || opChoiceInt < 1 || opChoiceInt > 5 {
			fmt.Println("Invalid choice! Please enter a number between 1 and 5.")
			flag++
			if flag > 3 {
				fmt.Println("Too many invalid inputs!\nTerminating Calculator!...")
				os.Exit(0)
			}
			continue
		}

		if opChoiceInt == 5 {
			fmt.Println("Terminating calculator!...Goodbye!")
			os.Exit(0)
		}

		number1, number2 = utils.TakeNumbers(result, checker)

		switch opChoiceInt {
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
		fmt.Println("The result is : ", result)
		checker++
	}
}
