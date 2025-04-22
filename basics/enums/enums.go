package main

import (
	"fmt"
	
)

type OrderStatus int
const (
	Received OrderStatus =iota
	Confirmed 
	Prepared
)
//enumerated types
func changeOrderStatus(status OrderStatus){
	fmt.Println("changing order status to",status)
}
func main() {
	changeOrderStatus(Prepared)
	fmt.Println("Shortcut triggered! 🟢")
}