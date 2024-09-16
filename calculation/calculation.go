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
