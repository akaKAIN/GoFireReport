package main

import "sync"

type AllElements struct {
	List []Element
	m 	*sync.Mutex
}

type Element struct {
	Name       string
	Id         string
	ElemType   interface{}
	ErrorMsg   string
	Signal     string
	IsSelected bool
	Handler    func(interface{}) interface{}
}
