package main

import "fmt"

type Context struct {
	state State
}

func (c *Context) do() {
	c.state.do()
}

type State interface {
	changeState(context *Context)
	do()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) changeState(context *Context) {
	fmt.Println("A -> B")
	context.state = &ConcreteStateB{}
}

func (s *ConcreteStateA) do() {
	fmt.Println("jump")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) changeState(context *Context) {
	fmt.Println("B -> A")
	context.state = &ConcreteStateA{}
}

func (s *ConcreteStateB) do() {
	fmt.Println("fly")
}

func main() {
	context := &Context{state: &ConcreteStateA{}}
	context.do()

	context.state.changeState(context)
	context.do()

	context.state.changeState(context)
	context.do()
}
