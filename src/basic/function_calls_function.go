package basic

import (

)

var a string

func abc() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "0"
	print(a)
	f2()
}

func f2() {
	print(a)
}

