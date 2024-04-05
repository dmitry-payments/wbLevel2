package main

import "fmt"

type PersonBuilder interface {
	SetName(name string) PersonBuilder
	SetAge(age int) PersonBuilder
	Build() Person
}

type Person struct {
	Name string
	Age  int
}

func NewPersonBuilder() PersonBuilder {
	return &ConcretePersonBuilder{}
}

type ConcretePersonBuilder struct {
	person Person
}

func (pb *ConcretePersonBuilder) SetName(name string) PersonBuilder {
	pb.person.Name = name
	return pb
}

func (pb *ConcretePersonBuilder) SetAge(age int) PersonBuilder {
	pb.person.Age = age
	return pb
}

func (pb *ConcretePersonBuilder) Build() Person {
	return pb.person
}

func getBobBuilder() PersonBuilder {
	return NewPersonBuilder().SetName("Bob").SetAge(30)
}

func main() {
	p := getBobBuilder().Build()

	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
