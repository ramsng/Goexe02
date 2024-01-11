package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := int(rand.Intn(250))
	fmt.Println("Random number is :", x)
	fmt.Println("which lies between")
	switch {
	case (x <= 100):
		for i := 0; i <= 100; i++ {
			fmt.Println(i)
		}
	case (x <= 200):
		for i := 100; i <= 200; i++ {
			fmt.Println(i)
		}
	case (x <= 250):
		for i := 200; i <= 250; i++ {
			fmt.Println(i)
		}
	default:
		{
			fmt.Println("Raqndom value out of range")
		}

	}
}
