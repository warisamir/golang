package main
import (
	"fmt"
)
func main(){
	// var nums[4]int
	// nums[0]=12
	var ar= make([]int,4)
	 nm:=[]int{}
	 nm=append(nm,0,4)
	//for increase predefined capacity
	ar=append(ar,1)
	ar=append(ar,2)
	nums:=[2][2]int{{3,5},{5,6}}
	m:=map[string]int {"price":300,"location":1500}
	m["aman"]=1
	fmt.Print(m,m["aman"]);
	fmt.Println(nm,cap(ar),nums[0:],ar)
}