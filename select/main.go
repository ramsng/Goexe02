package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("\niteration no:\t%d", i)
		channel1 := make(chan int)
		channel2 := make(chan int)

		time1 := time.Duration(10)
		time2 := time.Duration(10)

		go func() {
			time.Sleep(time1 * time.Millisecond)
			channel1 <- 01
		}()
		go func() {
			time.Sleep(time2 * time.Millisecond)
			channel2 <- 02
		}()
		select {
		case v1 := <-channel1:
			fmt.Printf("\n channel # : %v", v1)
		case v2 := <-channel2:
			fmt.Printf("\n channel # : %v", v2)
		}
	}
}
