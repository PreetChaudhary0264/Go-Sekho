package main

import "fmt"

func sum(result chan int, a, b int) {
	ans := a + b
	result <- ans
}

func main1() {
	result := make(chan int)   //aise channel bnate hai to usko unBuffered channel bolte hai....ye blocking hote hai

	go sum(result, 2, 3)

	ans := <-result
	fmt.Println(ans)
	//yha time.sleep ki zarurat nhi pdi bcoz channels me bhejna and receive krne blocking hota hai
	//ye partially sach hai bcoz ek orr type ke channels hote hai(buffered channel) aage pdenge
}