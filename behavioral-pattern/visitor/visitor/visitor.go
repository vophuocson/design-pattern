package visitor

import (
	"design-pattern-practice/behavioral-pattern/visitor/element"
	"fmt"
)

type VisitorType string

const (
	ZoomInType  VisitorType = "zoom in"
	ZoomOutType VisitorType = "zoom out"
)

type ZoomOut struct{}

func (v *ZoomOut) VisitForSquare(s *element.Square) {
	s.Side += 1
	fmt.Println(s.Side)
}

func (v *ZoomOut) VisitForCircle(s *element.Circle) {
	s.Side += 1
	fmt.Println(s.Side)
}

func (v *ZoomOut) VisitForRectangle(s *element.Rectangle) {
	s.Side += 1
	fmt.Println(s.Side)
}

type ZoomIn struct{}

func (v *ZoomIn) VisitForSquare(s *element.Square) {
	s.Side -= 1
	fmt.Println(s.Side)
}

func (v *ZoomIn) VisitForCircle(s *element.Circle) {
	s.Side -= 1
	fmt.Println(s.Side)
}

func (v *ZoomIn) VisitForRectangle(s *element.Rectangle) {
	s.Side -= 1
	fmt.Println(s.Side)
}

func CreateVisitor(visitorType VisitorType) element.Visitor {
	switch visitorType {
	case ZoomInType:
		return &ZoomIn{}
	case ZoomOutType:
		return &ZoomOut{}
	default:
		return nil
	}
}
