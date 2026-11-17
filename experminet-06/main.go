
package main

import (
	"fmt"
	"time"
)

// Result stores the output of each goroutine
type Result struct {
	Task   string
	Output string
}

// Task 1: Calculate square
func calculateSquare(ch chan<- Result) {
	fmt.Println("Square goroutine started")

	number := 5
	time.Sleep(2 * time.Second)

	square := number * number

	fmt.Println("Square goroutine completed")

	ch <- Result{
		Task:   "Square",
		Output: fmt.Sprintf("Square of %d = %d", number, square),
	}
}

// Task 2: Calculate cube
func calculateCube(ch chan<- Result) {
	fmt.Println("Cube goroutine started")

	number := 3
	time.Sleep(1 * time.Second)

	cube := number * number * number

	fmt.Println("Cube goroutine completed")

	ch <- Result{
		Task:   "Cube",
		Output: fmt.Sprintf("Cube of %d = %d", number, cube),
	}
}

// Task 3: Generate Fibonacci series
func calculateFibonacci(ch chan<- Result) {
	fmt.Println("Fibonacci goroutine started")

	n := 8
	a, b := 0, 1
	series := ""

	for i := 0; i < n; i++ {
		series += fmt.Sprintf("%d ", a)
		a, b = b, a+b
	}

	time.Sleep(3 * time.Second)

	fmt.Println("Fibonacci goroutine completed")

	ch <- Result{
		Task:   "Fibonacci",
		Output: fmt.Sprintf("Fibonacci series: %s", series),
	}
}

func main() {
	fmt.Println("===== CONCURRENT TASK EXECUTION =====")

	// Create a buffered channel
	ch := make(chan Result, 3)

	// Launch three goroutines
	go calculateSquare(ch)
	go calculateCube(ch)
	go calculateFibonacci(ch)

	fmt.Println("\nWaiting for all goroutines to finish...\n")

	// Collect results from all goroutines
	fmt.Println("----- Task Results -----")

	for i := 0; i < 3; i++ {
		result := <-ch
		fmt.Printf("%s: %s\n", result.Task, result.Output)
	}

	close(ch)

	fmt.Println("\nAll goroutines completed successfully!")
}
