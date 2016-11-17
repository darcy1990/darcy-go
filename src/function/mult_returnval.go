package function

import (

)

func MultReturnvalUnnamed(a int, b int) (int, int, int) {
	return a + b, a * b, a -b
}

func MultReturnvalNamed(a int, b int) (sum int, prod int, diff int) {
	sum = a + b
	prod = a * b
	diff = a -b
	return
}


