package main

import "fmt"

type Request struct {
	kind string
}

type Handler interface {
	handleRequest(request *Request)
	setNext(handler Handler)
}

type ConcreteHandler struct {
	next Handler
}

func (c *ConcreteHandler) handleRequest(request *Request) {
	fmt.Println("ConcreteHandler обрабатывает заявку типа:", request.kind)
	if c.next != nil {
		c.next.handleRequest(request)
	}
}

func (c *ConcreteHandler) setNext(handler Handler) {
	c.next = handler
}

func main() {
	handler1 := &ConcreteHandler{}
	handler2 := &ConcreteHandler{}
	handler3 := &ConcreteHandler{}

	handler1.setNext(handler2)
	handler2.setNext(handler3)

	request1 := &Request{kind: "Тип1"}

	handler1.handleRequest(request1)
}
