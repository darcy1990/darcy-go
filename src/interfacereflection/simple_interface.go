package interfacereflection

import (

)

type Simpler interface {
	Get() int
	Set(a int)
}

type Simple struct {
	val int
}

func (s *Simple) Get() int {
	return s.val
}

func (s *Simple) Set(a int) {
	s.val = a
}

func Call(s *Simple) {
	s.Set(20)
	println(s.Get())
}


