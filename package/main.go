package main

import (
	"fmt"
	"github.com/warisamir/golang/user"
	"github.com/warisamir/golang/auth"
)

func main() {
	auth.LoginwithCred("war","waris1918")
	fmt.Println("Shortcut triggered! 🟢")
	session:=auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email: "User@gmail.com",
		Name:  "john wick",
	}
	fmt.Println(user.Email,user.Name)
} 