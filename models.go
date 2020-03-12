package main

type Element struct {
	Name       string
	Id         string
	ElemType   interface{}
	ErrorMsg   string
	Signal     string
	IsSelected bool
	Handler    func(interface{}) interface{}
}
