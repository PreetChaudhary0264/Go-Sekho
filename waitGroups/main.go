package main

import (
	"fmt"
	"sync"
)

func dotask(id int, wg *sync.WaitGroup) {
	defer wg.Done()          //defer is line ko function me sbse baad me execute krega 
	                         //aise smjhlo ki wg.Done() decrement krta hai value by 1
	fmt.Println("doing task with task id",id)
}

func main(){
	var wg sync.WaitGroup

	for i := 0 ; i < 10; i++ {
		wg.Add(1)
		go dotask(i,&wg)
	}

	wg.Wait()  //ye main thread ko rok rha hai execute hone se
	//waitgroup me bs 3 hi chize yaad krni hai ADD, DONE, WAIT
}