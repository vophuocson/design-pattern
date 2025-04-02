package visitor

import (
	"design-pattern-practice/visitor/element"
	"design-pattern-practice/visitor/visitor"
)

func main() {
	shape := element.CreateShape(element.CircleType, 10)
	zoomOut := visitor.CreateVisitor(visitor.ZoomOutType)
	shape.Accept(zoomOut)
}
