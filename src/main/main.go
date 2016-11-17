package main

import (
	"fmt"
	"function"
	"arr"
)


func main() {
	fv := func() {
		println("Hello World")
	}
	fv()
	fmt.Printf("%T",fv)
	
	for i := 0; i < 10; i++ {
		fact := function.Factorial(i)
		fmt.Printf("%d ", fact)
	}
	
	println("")
	var sli = []int {1, 2, 3}
	println(arr.MinSlice(sli))
	println(arr.MaxSlice(sli))
}

