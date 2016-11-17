package arr

import (

)

func MinSlice(sli []int) int {
	min := sli[0]
	for _, val := range sli {
		if val < min {
			min = val
		}
	}
	
	return min
}

func MaxSlice(sli []int) int {
	max := sli[0]
	for _, val := range sli {
		if val > max {
			max = val
		}
	}
	
	return max
}

