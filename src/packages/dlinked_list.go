package packages

import (
	"container/list"
	"fmt"
)

func DlinkedList() {
	l := list.New()
	l.PushFront(101)
	l.PushFront(102)
	l.PushFront(103)
	
	for e:= l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d \n", e.Value)
	}
	
}

