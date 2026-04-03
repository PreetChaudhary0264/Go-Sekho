package main

import (
	"fmt"
	"time"
)
//hum safety bhi increase kr skte hai jaise agar hum chate hai ki koi iss func me sirf channel se values receive kr paye bhej nhi paye to iska syntax ye hoga result <-chan int
//ya agar sirf send kr paye receive nhi to fir => done chan<- bool
func queueSystem(queue <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	for email := range queue {
		fmt.Println("Sending email to ", email)
		time.Sleep(time.Second)
	}
}

func main23(){
	queue := make(chan string, 20)  //yha pe initial size dene se ye buffered channel bn jata hai,  agar is size se jyada values hai channel me tab wo blocking ki tarah kam krega 
	done := make(chan bool)
	go queueSystem(queue, done)

	for i := 1; i <= 10; i++ {
		queue <- fmt.Sprintf(" %d@gmail.com", i)
	}
    fmt.Println("Done sending")

	close(queue)  //agar channel close nhi kra to bhi deadlock aayga bcoz range expect krega ki channel me koi orr values aayngi and ye niche wali line hmesha wait krti rhe gi to deadlock aa gya isliye channel close krne bhot zarurui hai buffered me (waise dono channel me buffered ya unbuffered)
	<- done
}