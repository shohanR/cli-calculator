// TakeNumbers takes input from the user
// [Q]: why are you passing "num" as a parameter? What's the purpose?
// FIXME: this function has nothing to do with the calculation. Could be moved in another pkg (something like "utils") or put as a function in the main.go file.
// BUG: handle conversion between strings and float64 values

package calculation

// Add performs addition
func Add(num1, num2 float64) float64 {
	return num1 + num2
}

// Sub performs subtraction
func Sub(num1, num2 float64) float64 {
	return num1 - num2
}

// Mul performs multiplication
func Mul(num1, num2 float64) float64 {
	return num1 * num2
}

// Div performs division
func Div(num1, num2 float64) float64 {
	if num2 == 0 {
		return 0
	}
	return num1 / num2
}
