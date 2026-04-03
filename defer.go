package main

import "fmt"

func task(done chan bool) {
	defer func(){ done <- true}()   //agar maanlo processing me kuch error aagya to niche jayga hi nhi isliye isko defer function bnake uper likh diya
	fmt.Println("Processing")
}

func main3() {
	done := make(chan bool)

	go task(done)

	<-done
}