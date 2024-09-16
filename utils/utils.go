package utils

import (
	"fmt"
	"os"
	"strconv"
)

// takeNumbers is a utility function to take input numbers from the user.
// BUG: os.Exit(0) is wrong due to the following:
// 1. "0" means success. You're using it in a failure scenario.
// 2. It's the caller of this function that should handle the error returned by this function. It's
// not is duty to halt the program.
func TakeNumbers(result float64, checker int) (float64, float64) {
	var num1str, num2str string
	var num1, num2 float64

	if result == 0 || checker == 0 {
		fmt.Print("Enter first number: ")
		fmt.Scan(&num1str)
		var err error
		num1, err = strconv.ParseFloat(num1str, 64)
		if err != nil {
			fmt.Println("Error! Invalid input")
			os.Exit(0)
		}

		fmt.Print("Enter second number: ")
		fmt.Scan(&num2str)
		num2, err = strconv.ParseFloat(num2str, 64)
		if err != nil {
			fmt.Println("Invalid input!")
			os.Exit(0)
		}
	} else {
		num1 = result
		fmt.Print("Enter next number: ")
		fmt.Scan(&num2str)
		var err error
		num2, err = strconv.ParseFloat(num2str, 64)
		if err != nil {
			fmt.Println("Error! Invalid input!")
			os.Exit(0)
		}
	}
	return num1, num2
}
