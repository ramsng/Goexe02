package main

import (
	"fmt"
)

func main() {
	/* exercise : 24
	for i := 0; i <= 100; i++ {
		fmt.Println("Iteration:", i)
		x := rand.Intn(10)
		y := rand.Intn(10)
		switch {
		case x < 4 && y < 4:
			fmt.Println("Input 1 < 4 & Input 2 < 4 :", x, y)
		case x > 6 && y > 6:
			fmt.Println("Input 1 > 4 & Input 2 > 4 :", x, y)
		case x >= 4 && x <= 6:
			fmt.Println("Input 1 between 4 & 6 :", x, y)
		case y != 5:
			fmt.Println("Input 2 not 5 : ", x, y)
		default:
			fmt.Println("Input 1 & Input 2 is unique : ", x, y)
		}
	} */
	// exercise : 25
	/*
		for i := 0; i <= 42; i++ {
			fmt.Println("Iteration:", i)
			x := int(rand.Intn(5))
			switch x {
			case 0:
				fmt.Println("Lucky number is ZERO - ", x)
			case 1:
				fmt.Println("Lucky number is ONE - ", x)
			case 2:
				fmt.Println("Lucky number is TWO - ", x)
			case 3:
				fmt.Println("Lucky number is THREE - ", x)
			case 4:
				fmt.Println("Lucky number is FOUR - ", x)
			case 5:
				fmt.Println("Lucky number is FIVE - ", x)
			}
		}
	*/
	// exercise : 27
	/*
		i := 0
		for i <= 42 {
			fmt.Println("Iteration:", i)
			x := int(rand.Intn(5))
			switch x {
			case 0:
				fmt.Println("Lucky number is ZERO - ", x)
			case 1:
				fmt.Println("Lucky number is ONE - ", x)
			case 2:
				fmt.Println("Lucky number is TWO - ", x)
			case 3:
				fmt.Println("Lucky number is THREE - ", x)
			case 4:
				fmt.Println("Lucky number is FOUR - ", x)
			case 5:
				fmt.Println("Lucky number is FIVE - ", x)
			}
			i = i + 1
		}
	*/
	// exercise : 28
	/*
		i := 0
		for {
			if i > 100 {
				break
			}
			fmt.Println("Iteration:", i)
			x := int(rand.Intn(5))
			switch x {
			case 0:
				fmt.Println("Lucky number is ZERO - ", x)
			case 1:
				fmt.Println("Lucky number is ONE - ", x)
			case 2:
				fmt.Println("Lucky number is TWO - ", x)
			case 3:
				fmt.Println("Lucky number is THREE - ", x)
			case 4:
				fmt.Println("Lucky number is FOUR - ", x)
			case 5:
				fmt.Println("Lucky number is FIVE - ", x)
			}
			i = i + 1
		}
	*/
	// exercise : 29
	/*
		fmt.Println("EVEN NUMBERS BETWEEN 100 & 150")
		for i := 100; i >= 100; i++ {
			if i > 150 {
				break
			}
			if i%2 == 1 {
				continue
			} else {
				fmt.Printf("\n%d", i)
			}
		}
	*/
	// exercise : 29
	/*
		xi := []int{42, 43, 44, 45, 46, 47}
		for a, b := range xi {
			fmt.Printf("\noccurance : %d\tvalue : %d", a, b)
		}
	*/
	// exercise : 30
	/*
		x := map[string]int{
			"James":      42,
			"Moneypenny": 32,
		}
		for a, b := range x {
			fmt.Printf("\nRoll #:%d\tFisrt Name :%s", b, a)
		}
	*/
	/* exercise 31
		m := map[string]int{
			"James": 16,
			"Jimmy": 10,
			"Kim":   15,
		}
		for a, _ := range m {
			fmt.Printf("\n%s", a)
		}

		age := m["Kim"]
		fmt.Println("The age of Kim :", age)

		if v, ok := m["Kim"]; ok {
			fmt.Println("Kim record is with registry", v)
		}

		age = m["Rams"]
		fmt.Println("The age of Rams :", age)

		if v, ok := m["Rams"]; ok {
			fmt.Println("Rams record is not in the entry", v)
		}

		age = m["James"]
		fmt.Println("The age of James :", age)

		if v, ok := m["james"]; ok {
			fmt.Println("James record is present in registry", v)
		} else {
			fmt.Println("James record is not present in registry", v)
		}

		age = m["Jimmy"]
		fmt.Println("The age of Jimmy :", age)
	}
	*/
	// exercise 33
	/*
		for i := 0; i <= 100; i++ {
			if y := rand.Intn(5); y == 3 {
				fmt.Println("Random value is 3 ", y)
			} else {
				fmt.Println("Random value is not 3 ", y)
			}
		}
	*/
	fmt.Println("Values true (or) False :\t", true && true)
	fmt.Println("Values true (or) False :\t", true && false)
	fmt.Println("Values true (or) False :\t", true || true)
	fmt.Println("Values true (or) False :\t", true || false)
	fmt.Println("Values true (or) False :\t", false)

}
