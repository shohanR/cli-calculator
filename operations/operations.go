package operations

import (
	"cli-calculator/calculation"
	"cli-calculator/utils"
	"fmt"
)

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
				} else {
					result = calculation.Div(number1, number2)
				}
			}

			// FIXME: move this check away from here.
			history = OperationResult(opChoiceInt, number1, number2, result, operatorSign)
			return result, zeroDev, history
		}
	}
}

func OperationResult(opChoice int, number1, number2, result float64, operatorSign string) string {
	var resultHistory string
	if opChoice == 4 && number2 == 0 {
		resultHistory = "undefined!"
	} else {
		resultHistory = fmt.Sprintf("%v", number1) + operatorSign + fmt.Sprintf("%v", number2) + " = " + fmt.Sprintf("%v", result)
	}
	return resultHistory
}
