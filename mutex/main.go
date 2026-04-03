package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex   //mutes use kro struct me achi practice, wrna global bhi use kr skte thee in main function
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func(){
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1
	// p.mu.Unlock()   best practice hai ki ise bhi defer func me daaldo


}

func main() {
	var wg sync.WaitGroup
	post := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go post.inc(&wg)
		//agar hum yha wg.done() kr dete to wo ekdum execute ho jata bcoz upr wala operation blocking nhi hai
	}
    
	wg.Wait()
	fmt.Println(post.views)
}