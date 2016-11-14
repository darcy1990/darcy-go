package hello

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	fmt.Println(c)
}
