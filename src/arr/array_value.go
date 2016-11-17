package arr

import (

)

func ArrayValue() {
	var arr1 [3]int
	arr2 := arr1
	arr2[2] = 3
	
	for i := range arr1 {
		println(arr1[i])
	}
	
	for i := range arr2 {
		println(arr2[i])
	}
	
}

