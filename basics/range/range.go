package main
import (
	"fmt"
)
func main(){
	// nums:=[]int{5,6,7,8}
	// sum:=0
	// for i, num:=range nums{
	// 	sum=sum+i
	// 	fmt.Println(num,sum)
	// }
	// m:=map[string]string{"fname":"waris","lname":"amir"}
	// for _ ,v:=range m{
	// 	fmt.Println(v);
	// }
	for i,c:= range "golang"{
		fmt.Println(i,string(c));
	}
	 fn:=processIt()
	 fn(56)
}
func processIt() func (a int )int {
	return func(a int) int {
		return 5
	}
}