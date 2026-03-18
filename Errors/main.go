package main

import (
	// "errors"
	"fmt"
)

type myError struct {
	msg string
}

func (e myError) Error() string{
	return e.msg
}

func divide(a, b int) (int, error) {
	if b == 0 {
		// return 0, errors.New("Dividing by zero")
		// return 0, fmt.Errorf("Dividing by zero")
		return 0, myError{msg:"Dividing by zero"}
	}

	return a / b, nil
}

func main(){
	val,err := divide(2,0)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}