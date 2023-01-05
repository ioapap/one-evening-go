package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}

func Divide(x float64, y float64) (float64, error) {
	if x != 0 && y != 0 {
		return x / y, nil
	} else {
		return 0, errors.New("not possible")
	}
}
