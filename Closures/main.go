package main

import "fmt"

func process() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}
func main() {
	increment := process()
	fmt.Println(increment())
	fmt.Println(increment())
}