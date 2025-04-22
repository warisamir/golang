package main

import (
	"fmt"
	// "strconv"
)

// func add(a , b int)(int,int) {
// 	a=a+b
// 	b=a-b
// 	a=a-b
// 	return a, b
// }
// func getLagn(a,b string)(string,bool){
//  if _, err := strconv.Atoi(b); err == nil {
// 		return a, true
// 	}
// 	return "", false
// }
func processIt(fn func (a int) int){
	res:=fn(1)
	fmt.Println(res)
}
// func processIt() func (a int )int {
// 	return func (a int) int{
// 		return  4
// 	}
// }
func main() {
	// res2,res3:= getLagn("war","4")
	// res,res1 := add(3, 4)
	// fmt.Println(res,res1,res3,res2)
	fn:=func(a int) int {
		return 2
	}
	processIt(fn)
}
