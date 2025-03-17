package externalproduct

import "fmt"

type Drawing interface {
	Draw()
}

type Circle struct {
}

func (d *Circle) Draw() {
	fmt.Println("this is Draw() function of Circle belong external product")
}

type Square struct {
}

func (d *Square) Draw() {
	fmt.Println("this is Draw() function of Square belong external product")
}
