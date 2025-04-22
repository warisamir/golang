package main

import (
	"fmt"
)
func main(){

 whoAmi := func(i interface{}){
	switch c:=i.(type){
 case int:
	fmt.Println("its integer")
 case string:
	fmt.Println("it is string")
 case bool:
	fmt.Println("it is boolean ",c)
 }
 } 
 fmt.Println();
 whoAmi(55.6)
}