package main

import (
	"fmt"
	"maps"
)

func main(){
	m := make(map[string]string)
	fmt.Println(m)

	m["name"] = "preet"
	m["city"] = "gzb"
	fmt.Println(m["name"], m["phone"])  //agar koi key map me nhi hai to zero value return hoti hai

	fmt.Println(len(m))

    delete(m, "city")  //key delete krni hai
	fmt.Println(m)

	// clear(m)  pura map hi clear krna hai
	// copy(m2,m)  copy sirf slice prr hi kaam krta hai 

	m2 := map[string]int{"price":40, "age": 21}   //initialize map
	fmt.Println(m2)

	v, ok := m2["age"]  //basically v ke pass value aa jaygi agar map me age key present hai to wrna zero value aa jaygi and ok ke pass true/false rhega ki value present hai ya nhi
	fmt.Println(v)
	fmt.Println(ok)

	if ok {
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}

	//check maps are equal or not
     m3 := map[string]int {"name":3,"a":7}
	//  fmt.Println(m3 == m2)   not valid because wo objects hai
	fmt.Println(maps.Equal(m3,m2))   

}