package visitor

import (
	"design-pattern-practice/behavioral-pattern/visitor/element"
	"design-pattern-practice/behavioral-pattern/visitor/visitor"
)

func main() {
	shape := element.CreateShape(element.CircleType, 10)
	zoomOut := visitor.CreateVisitor(visitor.ZoomOutType)
	shape.Accept(zoomOut)
}
