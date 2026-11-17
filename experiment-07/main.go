
package main

import (
	"fmt"
	"math"
	"strings"
)

// Function 1: Division with error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

// Function 2: Square root with error handling
func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf(
			"cannot calculate square root of a negative number",
		)
	}

	return math.Sqrt(number), nil
}

// Function 3: Validate input
func validateInput(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("input cannot be empty")
	}

	return input, nil
}

func main() {
	fmt.Println("===== ERROR HANDLING IN GO =====")

	// Test division
	fmt.Println("\n1. Division Operation:")

	result, err := divide(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// Division by zero
	_, err = divide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
	}

	// Test square root
	fmt.Println("\n2. Square Root Operation:")

	root, err := squareRoot(25)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root of 25 =", root)
	}

	// Negative square root
	_, err = squareRoot(-9)

	if err != nil {
		fmt.Println("Error:", err)
	}

	// Test input validation
	fmt.Println("\n3. Input Validation:")

	input, err := validateInput("Hello Go")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Valid input:", input)
	}

	// Empty input
	_, err = validateInput("")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nProgram completed successfully!")
}
