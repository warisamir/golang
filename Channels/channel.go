package main

import (
	"math/rand"
	"fmt"
	"time"
)

func processNum(numChan chan int){
	for num:=range numChan{
		fmt.Println("proccessing number",num)
		time.Sleep(time.Second)
	}
}
func main() {
	numChan:=make(chan int)
	go processNum(numChan)
	for {
	numChan <- rand.Intn(100)
	}
	time.Sleep(time.Second*2)
	// messageChan:= make(chan string)
	// messageChan <- "ping pong"//channels are blocking operation
	// msg:= <-messageChan
	// fmt.Println(msg)
	fmt.Println("Shortcut triggered! 🟢")
}