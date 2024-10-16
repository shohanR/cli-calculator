package main

import (
	"fmt"
	"os"

	"cli-calculator/history"
	"cli-calculator/operations"
	"cli-calculator/utils"
)

func main() {
	var validHistory []string
	fmt.Println("\nWelcome to the GO CLI Calculator. Following are the options:")
	for {
		operationChoice, err := utils.GetoperationInput()
		if err != nil {
			fmt.Println(err, " Please enter a number between 1 and 5")
			continue
		}
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
