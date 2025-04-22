package main

import (
	"fmt"
	"time"
	// "time"
)

// func processNum(numChan chan int){
// 	for num:=range numChan{
// 		fmt.Println("proccessing number",num)
// 		time.Sleep(time.Second)
// 	}
// }

// func sum(result chan int ,num1 int, num2 int){
// 	res:= num1 + num2
// 	result <- res
// }

// func task (done chan bool){
// 	defer func(){ done<- true}()
// 	fmt.Println("processing...")

// }
// func emailSending (emailChan chan string ,done chan bool ){
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("sending email to ", email)
// 		time.Sleep(time.Second)
// 	}
// }
func main() {
	// emailChan := make(chan string ,100)
	// done := make (chan bool)
	// go emailSending(emailChan , done)
	// for i :=0;i<100;i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com",i)
	// }
	// fmt.Println("Shortcut triggered! 🟢")
	// close(emailChan)
	// <- done 
	// done:= make(chan bool)
	// go task(done)
	// <- done
	// result := make(chan int)
	// go sum(result,4,5)
	// res := <-result
	// fmt.Println(res)
	// numChan:=make(chan int)
	// go processNum(numChan)
	// for {
	// numChan <- rand.Intn(100)
	// }
	// time.Sleep(time.Second*2)
	// messageChan:= make(chan string)
	// messageChan <- "ping pong"//channels are blocking operation
	// msg:= <-messageChan
	// fmt.Println(msg)
	chan1 :=make(chan int)
	chan2 :=make (chan string )
	go func(){
		chan1<- 10
	}()
	go func(){
		chan2 <- "ping"
	}()

	for i:=0;i<2;i++{
		select {
		case chan1Val:= <- chan1:
			fmt.Println("received data form chan1 ",chan1Val)
		case chan2Val:= <- chan2:
			fmt.Println("receivd data from chan2", chan2Val)
	}
}
}
func emailSender(emailChan <-chan string , done chan<- bool){
	defer func() { done<- true}()

	for email:=range  emailChan{
		fmt.Println("sending Mail",email)
		time.Sleep((time.Second))
	}
}