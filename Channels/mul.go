package main

import (
	"fmt"
	// "time"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 5
	}()

	go func() {
		chan2 <- "Pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received value from chan1 ", chan1Val)
		case chan2Val := <- chan2:
		    fmt.Println("received value from chan2 ", chan2Val)
		}
	}
	fmt.Println("fone")
	// time.Sleep(2 * time.Second)
}