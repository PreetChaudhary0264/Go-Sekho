package main

import (
	"fmt"
)

// Arrays example
func arraysExample() {
	var fruits [3]string // array with fixed size
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "cherry"
	fmt.Println("Array:", fruits)
	fmt.Println("Length:", len(fruits))

	// Initialize array
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Initialized array:", numbers)
}

// Slices example
func slicesExample() {
	// Slice from array
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Println("Slice from array:", slice)

	// Make slice
	slice2 := make([]int, 3, 5)
	slice2[0] = 100
	fmt.Println("Make slice:", slice2)
	fmt.Println("Length:", len(slice2), "Capacity:", cap(slice2))

	// Append to slice
	slice2 = append(slice2, 200, 300)
	fmt.Println("After append:", slice2)
}

// Maps example
func mapsExample() {
	// Declare and initialize map
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	fmt.Println("Map:", ages)

	// Add key
	ages["Charlie"] = 35
	fmt.Println("After add:", ages)

	// Delete key
	delete(ages, "Bob")
	fmt.Println("After delete:", ages)

	// Check if key exists
	if age, exists := ages["Alice"]; exists {
		fmt.Println("Alice's age:", age)
	}
}

func main18() {
	fmt.Println("=== Go Basics - Part 2 ===")
	arraysExample()
	fmt.Println()
	slicesExample()
	fmt.Println()
	mapsExample()
}
