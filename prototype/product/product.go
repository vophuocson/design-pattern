package product

import "fmt"

type Pen interface {
	Draw()
	Clone() Pen
}

type Pencil struct{}

func (p Pencil) Draw() {
	fmt.Println("Drawing with a pencil.")
}

func (p Pencil) Clone() Pen {
	return p
}

type BallPen struct{}

func (b BallPen) Draw() {
	fmt.Println("Drawing with a ballpoint pen.")
}

func (p BallPen) Clone() Pen {
	return p
}

type BrushPen struct{}

func (b BrushPen) Draw() {
	fmt.Println("Drawing with a brush pen.")
}

func (p BrushPen) Clone() Pen {
	return p
}
