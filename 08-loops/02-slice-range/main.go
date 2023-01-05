package main

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}

func Sum(numbers ...int) (sum int) {
	sum = 0
	for _, num := range numbers {
		sum += num
	}
	return sum 
}
