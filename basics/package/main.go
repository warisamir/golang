package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/warisamir/golang/auth"
	"github.com/warisamir/golang/user"
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
	c:=color.New(color.FgHiGreen).Add(color.BgHiYellow)
	c.Println(user.Email,user.Name)
} 