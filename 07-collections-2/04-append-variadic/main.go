package main

import "fmt"

func Merge(a []string, b []string) []string {
	sum := append(a[:], b[:]...)
	return sum 
}

func Remove(a []string, index int) []string {
	elementRemove := append(a[:index], a[index+1:]...)
	return elementRemove
}

func RemoveLast(a []string) []string {
	last := Remove(a, len(a)-1)
	return last
}

func main(){
	test := Remove([]string{"1", "2", "3", "4", "5", "6", "7", "8"}, 4)
	fmt.Println(test)
}