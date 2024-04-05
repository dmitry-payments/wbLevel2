package main

import "fmt"

type Visitor interface {
	VisitSquare(square *Square)
	VisitCircle(circle *Circle)
}

type Shape interface {
	Accept(v Visitor)
}

type Square struct {
	Size float64
}

type Circle struct {
	Radius float64
}

func (s *Square) Accept(v Visitor) {
	v.VisitSquare(s)
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

type AreaVisitor struct {
	TotalArea float64
}

func (av *AreaVisitor) VisitSquare(square *Square) {
	av.TotalArea += square.Size * square.Size
}

func (av *AreaVisitor) VisitCircle(circle *Circle) {
	av.TotalArea += 3.14 * circle.Radius * circle.Radius
}

type VertexVisitor struct {
	totalVertex int
}

func (v *VertexVisitor) VisitSquare(square *Square) {
	v.totalVertex += 4
}

func (v *VertexVisitor) VisitCircle(circle *Circle) {
	v.totalVertex += 0
}

func main() {
	shapes := []Shape{
		&Square{Size: 5},
		&Circle{Radius: 3},
		&Circle{Radius: 2},
	}

	areaVisitor := &AreaVisitor{}

	vertexVisitor := &VertexVisitor{}

	for _, shape := range shapes {
		shape.Accept(areaVisitor)
		shape.Accept(vertexVisitor)
	}

	fmt.Println("Total area:", areaVisitor.TotalArea)
	fmt.Println("Total vertex:", vertexVisitor.totalVertex)
}
