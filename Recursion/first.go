package main

import "fmt"

func fibo(n int) int {   //normal function
	if n < 2 {
		return n
	}

	return fibo(n-1) + fibo(n-2)
}

func main() {

	fmt.Println(fibo(6))

	var fib func(a int) int  //basically hum anonymous function prr bhi recursion lga skte hai but hme unhe explicitly declare krna pdega

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(6))
}