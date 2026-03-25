package main

import "fmt"

// func print[T any](arr []T) {
// 	for _, val := range arr {
// 		fmt.Println(val)
// 	}
// }

type stack[T any] struct {   //we can also use generics in struct
	elements []T
}

func print[T int | string](arr []T) {   //iska matlab ab ye sirf int or string lega
	for _, val := range arr {
		fmt.Println(val)
	}
}


func multipleTypes[T any, V string](arr []T, name V) {   //we can also pass multiple generics types
	for _, val := range arr {
		fmt.Println(val, name)
	}
}

func main() {
   print([]int {1,2,3})
   print([]string {"1","2","3"})
//    print([]bool{true,false})

   myStack := stack[string] {
       elements: []string{"hi","hello"},
   }

   fmt.Println(myStack)

   multipleTypes([]int {1,2,3},"preet")
}