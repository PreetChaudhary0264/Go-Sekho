package main

import (
	"fmt"
)

func sum(a int,b int) int {
	return a + b
}

func multiple() (string, string, bool) {  //returnning multiple values
	return "preet", "chaudhary", true
}

//functions in parameters
func processIt (fn func(a int) int) {
	fn(1)
}

//return function from function
func process () func(a int) int {
	return func(a int) int {
		return 1
	}
}

//variadic functions
func add(nums ...int) int {   //ye sirf int values lega, any type ke liye hum interface{} use kr skte hai
	total := 0                //ye javascript ke rest operator ki tarah hai
	
	for _,num := range nums {
		total += num
	}
	return total
}

func main(){
    ans := sum(2,3)
	fmt.Println(ans)

	name,surname,isvalid := multiple()
	fmt.Println(name,surname,isvalid)

	fn := func(a int)int{
		return 1
	}

	processIt(fn)
    
	f := process()
	fmt.Println(f(1))

	// res := add(1,2,3,4,5)
	// fmt.Println(res)

	slice := []int {1,2,3,4,5}
	res := add(slice...)   //basically ye javascript ke spread operator ki tarah hai
	fmt.Println(res)

}