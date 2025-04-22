package main

import (
	"fmt"
	"time"
)


func task(id int){
	fmt.Println("doing task", id);
}

func main() {
	for i:=1;i<=100;i++{
	//  go task(i)
	go func(i int){
		fmt.Println(i)
	}(i)
	} 
	time.Sleep(time.Second*2) 
	fmt.Println("Shortcut triggered! 🟢")
}