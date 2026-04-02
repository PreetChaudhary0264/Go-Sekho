package main

import (
	"fmt"
	"time"
)

func main(){
    messages := make(chan string)

	go func(){
		messages <- "ping"
		messages <- "ping2"
	}()

	go func(){
		msg := <- messages
		msg2 := <- messages
		fmt.Println(msg)
		fmt.Println(msg2)
	}()

	// close(messages)   channel ko close krne ke liye

	//hum channels pr range bhi lga skte hai
	// for i := range messages {  //yha deadlock isliye aara hai kyuki values consume kr le and abi bhi range kr rhe hai to range wait krega ki orr values aaye ya fir channel close ho and kuch bhi mhi ho 5ha to deadlock aayga
	// 	fmt.Println(i)
	// }


	//channel close krne ke baad bhi usme values ko receive kiya ja skta hai but agar add krnege to panic aayga

	time.Sleep(2 * time.Second)


	numC := make(chan int)

	numC <- 5

	o := <- numC
	fmt.Println(o)

}