// FIXME: handle the goodness of the user's values. "d" is not a valid input here.
// TODO: let's do only one operation at time. Every time you want to perform an operation asks for all the data you need. NO OPERATION ON THE PREVIOUS OPERATION'S RESULT.
package main

import (
	"cli-calculator/history"
	"cli-calculator/operations"
	"cli-calculator/utils"
	"fmt"
	"os"
)

// TODO: try to remove logic from the main() since it's too much bloated with logic, checks, and so on.
func main() {
	var validHistory []string
	fmt.Println("\nWelcome to the GO CLI Calculator. Following are the options:")
	for {
		operationChoice, err := utils.GetoperationInput()
		if err != nil {
			fmt.Println(err, " Please enter a number between 1 and 5")
			continue
		}
		// TODO: if no operations have been made, say something like "No operations done..." to the user.
		if operationChoice == 5 {
			history.HistoryPrinter(validHistory)
			fmt.Println("\n\nTerminating calculator!...Goodbye!")
			os.Exit(0)
		}
		result, zeroDev, tempHistory := operations.GetoperationCalculation(operationChoice)
		if zeroDev == 1 {
			fmt.Println("\nundefined!")
		} else {
			fmt.Println("\nThe result is : ", result)
			validHistory = append(validHistory, tempHistory)
		}
	}
}
