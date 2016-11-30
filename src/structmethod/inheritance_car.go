package structmethod

import (

)

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	wheelCount int
}

type Mercedes struct {
	Car
	
}

func (c *Car) NumberOfWheels() int {
	return c.wheelCount
}

func (c *Car) SetWheelCount(count int) {
	c.wheelCount = count
}

func (m *Mercedes) SayHiToMerkel() {
	println("Hi, I am Mercedes")
}



