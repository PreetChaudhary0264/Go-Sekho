package main

//packages

import (
	"github.com/PreetChaudhary0264/Go-Sekho/Packages/auth"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredential()

	user := auth.User{
		Email: "test@example.com",
	}
	// println(user.Email)
	color.Red(user.Email)
}