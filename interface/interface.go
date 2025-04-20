package main

import (
	"fmt"
)

// add er in the last 
type paymenter interface{
	pay(amount float32)
} 



type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32){
// razorpaymentlink:= razorpay{}
// stripepayment:= stripe{}
// razorpaymentlink.pay(amount)
p.gateway.pay(amount)
}
type razorpay struct {}
func  (r razorpay) pay(amount float32){
	fmt.Println("making payment using razorpay",amount)
 }
type stripe struct{}
func (s stripe) pay(amount float32){ 
	fmt.Println("making payment using stripe gateway",amount)
}

type fakepayment struct{}
func (f fakepayment) pay(amount float32){
	fmt.Println("making payment using fake payment gatway!!!");
}

func main() {
	fakepaymentgateway:= fakepayment{}
	// razorpaymentlink:= razorpay{} 
	newPayment:= payment{
		gateway: fakepaymentgateway,
	}
	newPayment.makePayment(100)
	fmt.Println("Shortcut triggered! 🟢")
}