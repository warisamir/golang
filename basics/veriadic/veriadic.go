package main

import "fmt"

func sum(nums ...int) int{
total:=0
for _,num:= range nums{
	total=total+num
	fmt.Println(total)
}
 return total
}
func main(){
	res:=sum(1,2,3,4,5)
	fmt.Println(res)
	fmt.Println(1,2,3,4,5, "hello")
}
