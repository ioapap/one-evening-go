package main

import (
	"fmt"
)

func main() {
	err := RunSafely(func() {
		Divide(10, 0)
	})
	fmt.Println(err)
}

func Divide(x, y float64) float64 {
	if y == 0 {
		panic("division by zero")
	}

	return x / y
}

func RunSafely(f func()) (err error) {
	// Define a function that handles panics and returns an error.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("a panic occurred: %v", r)
		}
	}()

	// Execute the function.
	f()
	return err
}

//This function defines and calls a deferred function that handles panics. If a panic occurs during the execution of f, the deferred function will recover from the panic and return an error indicating that a panic occurred. If no panic occurs, the deferred function will not be called and the RunSafely function will return nil. Finally, the RunSafely function calls f to execute it.
