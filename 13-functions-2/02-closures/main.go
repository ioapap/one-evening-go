package main

import "fmt"

func WordGenerator(words []string) func() string {
	currentIndex := 0 
	return func() string {
		word := words[currentIndex]
		currentIndex++
		if currentIndex >= len(words) {
			currentIndex = 0
		}
		return word
	}
}

func main() {
	continents := []string{
		"Africa",
		"Antarctica",
		"Asia",
		"Australia",
		"Europe",
		"North America",
		"South America",
	}

	generator := WordGenerator(continents)

	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}
}
