package main

import "fmt"

func main() {
	xi := []int{1, 61, 30, 51, 23}
	for j, w := range xi {
		fmt.Println("index and values :", j, w)
	}

	xj := map[string]int{
		"Rams":   24,
		"Jimmy ": 22,
		"Rajesh": 21,
	}

	for a, b := range xj {
		fmt.Println("CONDITIONALS VARIABLES :", a, b)
	}
}
