// FIXME: handle the goodness of the user's values. "d" is not a valid input here.
// TODO: let's do only one operation at time. Every time you want to perform an operation asks for all the data you need. NO OPERATION ON THE PREVIOUS OPERATION'S RESULT.
package main

import (
	"cli-calculator/calculation"
	"cli-calculator/utils"
	"fmt"
)

// TODO: try to remove logic from the main() since it's too much bloated with logic, checks, and so on.
func main() {
	var history []string
	// TODO: "checker" identifier is not meaningful. What's its purpose?
	fmt.Println("\nWelcome to the GO CLI Calculator. Following are the options:")
	for {
		var operationChoice, err = utils.GetoperationInput()
		if err != nil {
			fmt.Println(err, " Please enter a number between 1 and 5")
			continue
		}
		if operationChoice == 5 {
			utils.HistoryPrinter(history)
		}
		var result, zeroDev, tempHistory = GetoperationCalculation(operationChoice)
		if zeroDev == 1 {
			fmt.Println("\nThe result is :undefined!")
		} else {
			fmt.Println("\nThe result is : ", result)
		}
		history = append(history, tempHistory)
	}
}

func GetoperationCalculation(opChoiceInt int) (float64, int, string) {
	for {
		var number1, number2, result float64
		var err error = nil
		var zeroDev int = 0
		number1, number2, err = utils.TakeNumbers()

		if err != nil {
			fmt.Println(err, " Try again!")
		} else {
			var operatorSign, history string
			switch opChoiceInt {
			case 1:
				result = calculation.Add(number1, number2)
				operatorSign = " + "
			case 2:
				result = calculation.Sub(number1, number2)
				operatorSign = " - "
			case 3:
				result = calculation.Mul(number1, number2)
				operatorSign = " x "
			case 4:
				if number2 == 0 {
					fmt.Println("\n\nCannot divide by zero!")
					zeroDev = 1
					operatorSign = " / "
				}
				result = calculation.Div(number1, number2)
			}

			if opChoiceInt == 4 && number2 == 0 {
				history = fmt.Sprintf("%v", number1) + operatorSign + fmt.Sprintf("%v", number2) + " = undefined"
			} else {
				history = fmt.Sprintf("%v", number1) + operatorSign + fmt.Sprintf("%v", number2) + " = " + fmt.Sprintf("%v", result)
			}
			return result, zeroDev, history
		}
	}
}
