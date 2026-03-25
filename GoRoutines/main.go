package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for i := range 3 {
		fmt.Println(s,":", i)
	}
}

func main(){
	f("direct")

	go f("goroutine")  //basicalyy ab ye go wale function concurrently run krenge

	//goroutines on anonymous function
	go func(s string){
		fmt.Println(s)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}