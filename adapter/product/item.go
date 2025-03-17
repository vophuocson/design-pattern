package product

import (
	externalproduct "design-pattern-practice/adapter/external-product"
	"fmt"
)

type Item interface {
	Sketch()
}

type Rectangle struct {
}

func (i *Rectangle) Sketch() {
	fmt.Println("this is Sketch() function of Rectangle belong to product")
}

type Triangle struct {
}

func (i *Triangle) Sketch() {
	fmt.Println("this is Sketch() function of Triangle belong to product")
}

type Circle struct {
	Circle externalproduct.Drawing
}

func (i *Circle) Sketch() {
	i.Circle.Draw()
}

type Square struct {
	Square externalproduct.Drawing
}

func (d *Square) Sketch() {
	d.Square.Draw()
}
