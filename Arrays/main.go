package main

import "fmt"

func main(){

	// var nums[4]float64
	var nums[4]bool
	nums[0] = true
	fmt.Println(nums)

	// a := 5
	// var b float32 = float32(a)

	// same int(a)

	fmt.Println(len(nums))

	num := [3]int{1,2,3}
	fmt.Println(num)

	//2d array
	num2 := [2][2]int {{1,2},{3,4}}
	fmt.Println(num2)

	//if we want ki compiler khud size calculate kree 
	var f =  [...]int{1,2,3,4,5,6,7}
	fmt.Println(len(f))

}