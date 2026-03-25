package main 

import "fmt"

// type orderStatus int

// const (
// 	Received orderStatus = iota    //iota is an untyped integer, it is zero indexed     
// 	Processing                     //its value will be 1
// 	Shipped                        //its value will be 2
// 	Delivered
// 	Cancelled
// )

// func changeOrderStatus(status string){
// 	fmt.Println("changing status to", status)
// }

type orderStatus string

const (
	Received orderStatus = "received"
	Processing orderStatus = "processing"
	Cancelled orderStatus = "cancelled"
)

func changeOrderStatus(status orderStatus){
	fmt.Println("changing status to", status)
}


func main(){
	changeOrderStatus(Cancelled)
}