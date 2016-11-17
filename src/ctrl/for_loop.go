package control

import (
	"fmt"
)

func ForLoop() {
	for i := 1; i <= 15; i++ {
		fmt.Printf("%v\n", i)
	}
}

func Goto() {
	i := 0
	Here:
		fmt.Printf("%v\n", i)
		i++
		if i > 15 {
			return
		}
		goto Here
}

