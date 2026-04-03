package main

import "fmt"


func recoverExample(){

	defer func(){   //recover will only worki in defer functions
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic ", r)
		}
	}()

	fmt.Println("Before panic")
	panic("Something went wrong")
	fmt.Println("After panic")
}

func main() {
	recoverExample()
	fmt.Println("Over")
}