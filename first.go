package main

import "fmt"
import "time"


// a := 10 //global variable, ye error dega kyuki global variable ko hum aise declare nhi kr skte hai , shorthand syntax work nhi krta global me

func main() {
    fmt.Println("Hello World")

    var name string = "preet"
    fmt.Println("My name is", name)

    var age = 21  //ye khud type infer kr lega
    fmt.Println("My age is", age)

    //short hand declaration
    city := "ghaziabad"        
    fmt.Println("I live in", city)

    //agar hme variable declare krna haoi initilaize nhi to hum aise kr skte hai
    var country string

    country = "India"
    fmt.Println("I live in", country)

    //constants
    const pi = 3.14
    //kya constants ko shorthand syntax me declare kr skte hai? nahi kr skte hai
    //pi := 3.14 //ye error dega kyuki constants ko hum aise declare nhi kr skte hai , shorthand syntax work nhi krta constants me

    const (  //multiple constants ko ek sath declare krne ke liye
        a = 10
        b = "hi"
    )
    fmt.Println(a, b)

    //while loop using for loop
    // i := 1
    // for i <= 3 {
    //     fmt.Printf("%d",i)
    //     fmt.Printf("name: %s age: %d","Preet",21)
    //     i++
    // }

    //infinite loop
    // for {

    // }
    

    //for loop
    for i := 1; i < 3; i++ {

    }
    
    //range loop
    for j := range 3 {   //ye last range exclude hoti hai, to ye 0 se 2 tak loop chalayega
        fmt.Println(j)
    }

    //if else normal hote hai bs bina brackets ke likhte hai
    //hum if condition me bhi variable initilize kr skte hai and wo if ke ander bhi vlid hoga and next else if or else ke ander bhi valid hoga
    if age := 15; age > 18{
        fmt.Println("You are an adult")
    }else {
        fmt.Println("You are a minor")
    }

    //normal switch case

    k := 5
    switch k {
    case 1 :
        fmt.Println("One")
    case 2:
        fmt.Println("two")
    default:
        fmt.Println("other")
    }
    
    //multiple conditions in switch case
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend!")
    default:
        fmt.Println("It's a weekday.")
    }

    //type switch case
    whoAMI := func (i interface{}) {
        switch t := i.(type){     //switch i.(type){      ye bhi valid hai
        case int:
            fmt.Printf("I am an integer and my value is %d\n", t)
        case string:
            fmt.Printf("I am a string and my value is %s\n", t)
        default:
            fmt.Printf("I am of a different type: %T\n", t)
        }
    }

    whoAMI(42)
    
}