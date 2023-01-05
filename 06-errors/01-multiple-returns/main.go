package main

import "fmt"

func main() {
	x := "World!"
	y := "Hello,"

	x, y = Swap(x, y)

	fmt.Println(x, y)
}

func Swap(a string, b string) (string, string) {
	return b,a
}
