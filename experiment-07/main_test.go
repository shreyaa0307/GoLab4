
package main

import (
	"math"
	"testing"
)

// Test 1: Valid division
func TestDivideValid(t *testing.T) {
	result, err := divide(10, 2)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != 5 {
		t.Errorf("Expected 5, got %f", result)
	}
}

// Test 2: Division by zero
func TestDivideByZero(t *testing.T) {
	_, err := divide(10, 0)

	if err == nil {
		t.Error("Expected division by zero error")
	}
}

// Test 3: Valid square root
func TestSquareRootValid(t *testing.T) {
	result, err := squareRoot(25)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if math.Abs(result-5) > 0.000001 {
		t.Errorf("Expected 5, got %f", result)
	}
}

// Test 4: Negative square root
func TestSquareRootNegative(t *testing.T) {
	_, err := squareRoot(-9)

	if err == nil {
		t.Error("Expected negative square root error")
	}
}

// Test 5: Valid input
func TestValidateInputValid(t *testing.T) {
	result, err := validateInput("Hello Go")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != "Hello Go" {
		t.Errorf("Expected Hello Go, got %s", result)
	}
}

// Test 6: Empty input
func TestValidateInputEmpty(t *testing.T) {
	_, err := validateInput("")

	if err == nil {
		t.Error("Expected empty input error")
	}
}
