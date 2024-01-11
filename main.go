// ** CONDITIONAL INSTRUCTIONS WORK OUT **//
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Intn(2)
	b := rand.Intn(2)
	fmt.Printf("** Basic if function work out - Input-1: %d\tInput-2: %d\n", int(a), int(b))
	if a == b {
		fmt.Println("Random generation equal")
	} else if a > b {
		fmt.Println("Random  generation unequal", a, b)
	} else {
		fmt.Println("Random generation unequal", b, a)
	}

	a = rand.Intn(2)
	b = rand.Intn(2)
	fmt.Printf("** if function - declaration + logic - Input-1: %d\tInput-2: %d\n", int(a), int(b))
	if c := rand.Intn(2); a == b && a == c {
		fmt.Println("Random generation equal - Lucky most", a, b, c)
	} else if a == b || a == c || b == c {
		fmt.Println("Random generation equal - Lucky 2 combination", a, b, c)
	} else {
		fmt.Println("Random generation - No luck", a, b, c)
	}

	switch int(rand.Intn(2)) {
	case 0:
		fmt.Println("Male wins")
	case 1:
		fmt.Println("Female wins")
	}

	switch {
	case a > b:
		fmt.Println("Tough to play with gents", a, b)
	case a < b:
		fmt.Println("Tough to play with ladies", a, b)
		fallthrough
	case a > rand.Intn(2) && b > rand.Intn(2):
		fmt.Println("Tough to play with both Gent and Ladies", a, b)
	default:
		fmt.Println("Eazy to play with both Gents and Ladies", a, b)
	}
	// TRADITIONAL FOR LOOP
	for i := 1; i < 5; i++ {
		fmt.Printf("\nValue of i : %d on iteration:\t%d", i, i*10)
		for j := i * 100; j <= 200; j++ {
			fmt.Printf("\nValue of j %d on iteration : %d", j, j*10)
			j = j * 10
		}
	}
	// FOR LOOP WITH ZERO DECLARATION
	i := 10
	for i <= 100 {
		fmt.Printf("\nvalue of k: %d computing to value : %d ", i, i*10)
		i = i * 10
		i++
	}
	for {
		if i > 200 {
			break
		} else {
			i = i + rand.Intn(1000)
			fmt.Printf("\n Value of l: %d", i)
		}
	}
}
