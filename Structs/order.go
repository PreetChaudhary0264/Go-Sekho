package main

import (
	"fmt"
	"time"
)

type customer struct {
	name string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer              //struct embedding
}


//struct ke sath functions ko attach krna
func (o *order) changeStatus(status string){  //o ki jagah kuch bhi dedo but best practice is struct ka pehle charater do
    o.status = status
}

//receiving type
func (o order) getAmount() float32 {  //yha pointer ki zarurat nhi kyuki sirf get kr rhe hai
	return o.amount
}

//constructor alternative
func newOrder(id string, amount float32, status string) *order{
	myOrder := order{
		id: id,
		amount:amount,
		status: status,
		customer: customer{
			name: "preet",
			phone: "1234567890",
		},
	}
	return &myOrder
}



func main() {
    // var myOrder order = order{   //ye bhi valid hai
	// 	id: "123",
	// }
    
	//agar yha koi value miss kr dete hai deni to bydefault zero value set ho jati hai
	// myOrder := order{
	// 	id:"12345",
	// 	amount:50.0,
	// 	status:"received",
	// }

	language := struct {
		name string
		isGood bool
	}{"golang",true}

	fmt.Println(language)

	myOrder := newOrder("123",30.0,"paid")

	myOrder.createdAt = time.Now()
	fmt.Println(myOrder.status)
	fmt.Println(myOrder)

	myOrder.changeStatus("confirmed")
	fmt.Println(myOrder.status)
	
}