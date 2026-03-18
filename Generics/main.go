package main

import "fmt"

func print[T any](arr []T) {
	for _, val := range arr {
		fmt.Println(val)
	}
}

func main() {
   print([]int {1,2,3})
   print([]string {"1","2","3"})
}