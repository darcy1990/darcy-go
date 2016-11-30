package structmethod

import (
	
)

type Rectangle struct {
	Length int
	Width int
}

func Area(rt *Rectangle) int {
	return rt.Length * rt.Width

}

func Perimeter(rt *Rectangle) int {
	return (rt.Length + rt.Width) * 2
}

