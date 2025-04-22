package main

import (
	"fmt"
	
)

func closures() func() int{
   var ct int = 0
   return func() int {
	ct+=1
	return ct
   }
}
func main() {
	increment:= closures();
	fmt.Println(increment())
}