package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// takeNumbers is a utility function to take input numbers from the user.
// BUG: os.Exit(0) is wrong due to the following:
// 1. "0" means success. You're using it in a failure scenario.
// 2. It's the caller of this function that should handle the error returned by this function. It's
// not is duty to halt the program.
func TakeNumbers() (float64, float64, error) {
	var num1str, num2str string
	var num1, num2 float64

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1str)
	var err error
	num1, err = strconv.ParseFloat(num1str, 64)
	if err != nil {
		return 0, 0, errors.New("error! Invalid input for first number")
	}

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2str)
	num2, err = strconv.ParseFloat(num2str, 64)
	if err != nil {
		return 0, 0, errors.New("error! Invalid input for second number")
	}
	return num1, num2, nil
}

func GetoperationInput() (int, error) {
	var operationChoice string
	fmt.Print("\n1. Addition\n2. Subtraction\n3. Multiplication\n4. Division\n5. End\n")
	fmt.Print("Enter operation: ")
	fmt.Scan(&operationChoice)

	opChoiceInt, err := strconv.Atoi(operationChoice)
	if err != nil || opChoiceInt < 1 || opChoiceInt > 5 {
		return 0, errors.New("invalid input")
	}
	return opChoiceInt, nil
}

func HistoryPrinter(history []string) {
	if len(history) > 0 {
		for indexNum, indexVal := range history {
			fmt.Println(indexNum+1, ". ", indexVal)
		}
	}
	fmt.Println("\n\nTerminating calculator!...Goodbye!")
	os.Exit(0)
}
