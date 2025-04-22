package main

import (
	"fmt"
	"time"
)
type order struct {
	id string 
	amount float32
	status string
	createdAt time.Time
	customer
}
type customer struct{
	name string 
	phone string
}
// func newOrder(id string , amount float32,status  string) *order{
//  myOrder:=order{
// 	id: id,
// 	amount: amount,
// 	status: status,
// 	}
// 	return &myOrder
// }
// func (o *order) changeamount(amount float32){
// 	o.amount=amount;
// }
// func (o order) getAmount() float32{
// 	return o.amount
// }
func main() {
	newOrder:=order{
		id:"1",
		amount: 30.0,
		status: "Received",
		customer:customer{
				name:"jon",
				phone:"91305034",
			},
		}
	fmt.Println(newOrder.customer)
	lang:=struct{
		name string
		isGood bool
	}{"anam", false}
	fmt.Println(lang);
	// myOrder:=newOrder("123",20.50, "completed");
	// myOrder.changeamount(127.39)
	// fmt.Println(myOrder.amount)
	fmt.Println("Shortcut triggered! 🟢")
}