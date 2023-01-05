package main

import "fmt"

func main() {
	Greet("Hello, Alice")
	Greet("Hello, Bob")
}

func Greet(name string) {
	fmt.Println(name)
}
