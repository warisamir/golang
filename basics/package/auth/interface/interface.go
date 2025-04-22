package main

import (
	"fmt"
)

// add er in the last 
type paymenter interface{
	pay(amount float32)
	refund(amount float32,account string)
} 



type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32){
p.gateway.pay(amount)
}
func (p payment) refundinitiation(amount float32, account string){
	p.gateway.refund(amount,account)
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

type paypal struct{}
func (p paypal) pay(amount float32){
	fmt.Println("making payment using paypal payment gateway",amount);
}
func (p paypal) refund(amount float32 , account string){
	fmt.Printf("refund of %v has been initiated to accound-id %v \n",amount,account)
}
func main() {
	// fakepaymentgateway:= fakepayment{}
	// razorpaymentlink:= razorpay{} 
	paypalpaymentLink:= paypal{}
	newPayment:= payment{
		gateway: paypalpaymentLink,
	}
	newPayment.makePayment(100);
	newPayment.refundinitiation(400,"123453")
	fmt.Println("Shortcut triggered! 🟢")
}