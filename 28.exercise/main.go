package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}
func init() {
	fmt.Println("Numbers from 1 to 100")
}
