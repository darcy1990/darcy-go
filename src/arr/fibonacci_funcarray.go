package arr

import (

)

func FiboFunArray(n int) []int {
	var arr = make([]int, n)
	for i := range arr {
		arr[i] = fibo(i + 1)
	}
	return arr
}

func fibo(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibo(n -1) + fibo(n - 2)
}

