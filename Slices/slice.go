package main

import (
	"fmt"
	"slices"
)

func main(){

	var nums[]int      //uninitialized nil hoti hai
	fmt.Println(nums == nil)

	// nums = make([]int, 2)  //list ki initial length 2 ho gyi and wo value 0 hai ab kuch bhi append kroge to wo un 2 values ke baad me append hogi
	// fmt.Println(nums)

	nums = make([]int, 0, 5)  //ye 5 initial capacity hai khud resize hoti hai like arrayList in java
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	fmt.Println(nums)
	fmt.Println(cap(nums))
	nums = append(nums, 5)
	nums = append(nums, 6)
	fmt.Println(cap(nums))
	fmt.Println(nums)

	num := []int{}  //aise bhi slice ko declare kr skte hai
	fmt.Println(len(num))
	fmt.Println(cap(num))

	nums[3] = 10  //ye sirf value update krne ke loiye hai agar hme size badahane hai to append hi use krna 
	fmt.Println(nums)

	slice1 := []int {1,2,3}
	slice2 := make([]int, len(slice1))

	copy(slice2,slice1)  //ye slices ko copy krne ke liye use hota hai
	fmt.Println(slice1, slice2)

	//slice operator
	fmt.Println(slice2[:2])

	fmt.Println(slices.Equal(slice1,slice2))
    
	//2d slices
	// var slice2d = [][]int{{}}

	slice2d := make([][]int, 3)
	fmt.Println(slice2d)

	slice2d = append(slice2d, []int {1,2})
	fmt.Println(slice2d)
	for i := range 3 {
		slice2d[i] = make([]int, 4)
	}
	fmt.Println(slice2d)
}