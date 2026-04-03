package main

import (
	"fmt"
	"os"
)

func main() {
	f,err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	fileinfo,err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("fileName ",fileinfo.Name())
	fmt.Println("File size" , fileinfo.Size())
	fmt.Println(fileinfo.IsDir())
	fmt.Println("file Permission ",fileinfo.Mode())
	fmt.Println("fileName modified at",fileinfo.ModTime())

	buffer := make([]byte, fileinfo.Size())

	length,err := f.Read(buffer)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// println("data", length, buffer)
    println("length",length)
	for i := 0; i < len(buffer); i++ {
		print(string(buffer[i]))
	}
    
	println()
	filedata,err := os.ReadFile("example.txt")  //but ye readFile saari files ko ek sath load krta hai memory me to ye choti file ke liye shi hai bdi ke liye nhi resources kam pd skte hai
	if err != nil {
		panic(err)
	}
	println("data -> ", string(filedata))

	f2,err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}

	f2.WriteString("Hi go")
	f2.WriteString("ye append hoga replace nhi")
}