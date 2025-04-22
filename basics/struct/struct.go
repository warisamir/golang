package main

import (
	"fmt"
	"time"
)
// type order struct {
// 	id string 
// 	amount float32
// 	status string
// 	createdAt time.Time
// 	date time.Weekday
// }

func () {
	order := order{
		id:"1",
		amount :4000.0,
		status: "completed",
	}
	order.createdAt= time.Now()
	order.date = time.Now().Weekday()
	fmt.Print(order.createdAt,order.date);
}