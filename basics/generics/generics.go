package main

import (
	"fmt"
	
)

func printSlice[T comparable,V string](items []T, name V){
	for _,item := range items {
		fmt.Println(item,name)
	}
}
//LIFO

type stack[T any] struct {
elements []T
}
func main() {
	mystack := stack[int]{
		elements: []int {1,2,3},
	}
	mystacks:=stack[bool]{
		elements: []bool {true ,false},
	}
	fmt.Println(mystack,mystacks)
	nums:= []string{"golang", "typescript","javascript"}
	printSlice(nums, "john the done")
	// printSlice([]int{2,4,45,69})
	// fmt.Println("Shortcut triggered! 🟢")
}