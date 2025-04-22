package main

import (
	// "fmt"
	"github.com/fatih/color"
)

func main() {
	c:= color.New(color.BgHiWhite).Add(color.BlinkSlow)
	c.Println("welcome")
}