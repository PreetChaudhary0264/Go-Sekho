package main

import "fmt"

func main(){
	nums := []int {1,2,3,4,5}

	for i,num := range nums {   //i index hai and num value
		fmt.Println(i,num)
	}

	m := map[string]string {"name":"preet"}

	for k,v := range m {   //k = key hai and v = value
		fmt.Println(k,v)
	}

	for i,ch := range "golang" {
		fmt.Println(i,string(ch))   //agar strin gme convert nhi krnege to unicode aaynge
	}
}