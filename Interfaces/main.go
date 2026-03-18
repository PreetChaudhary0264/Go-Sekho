package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {  //ise krne se rectangle ne geometry ko implement kra hai
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {   //ise krne se circle ne geometry ko implement kra hai
	return math.Pi * c.radius * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func detectCircle(g geometry) {   //basically hu detect kr rhe hai ki geometry circle hai ya nhi
	if c, ok := g.(circle); ok {   //bina geometry ko implement kre ye error dega because us time circle geometry ko implement nhi kr rha hoga
		fmt.Println("Cicrle is of radius", c.radius)
	}
}

func main() {
    rect := rectangle {
		width: 3,
		height: 4,
	}
	circ := circle {
		radius: 2,
	}

	measure(rect)
	detectCircle(rect)  //yha hmne geometry rectangle pass kri hai to yha pe wo if ke andr nhi jayga
	detectCircle(circ)
}