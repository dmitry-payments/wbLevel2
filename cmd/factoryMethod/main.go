package main

import "fmt"

type Product interface {
	use() string
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) use() string {
	return "Используется продукт A"
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) use() string {
	return "Используется продукт B"
}

type Creator interface {
	createProduct() Product
}

type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) createProduct() Product {
	return &ConcreteProductA{}
}

type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) createProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	creatorA := &ConcreteCreatorA{}
	creatorB := &ConcreteCreatorB{}

	productA := creatorA.createProduct()
	productB := creatorB.createProduct()

	fmt.Println(productA.use())
	fmt.Println(productB.use())
}
