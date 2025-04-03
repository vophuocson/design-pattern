package element

type ShapeType string

const (
	SquareType    ShapeType = "square"
	CircleType    ShapeType = "circle"
	RectangleType ShapeType = "rectangle"
)

type Shape interface {
	Accept(v Visitor)
}

type Square struct {
	Side int
}

func (s *Square) Accept(v Visitor) {
	v.VisitForSquare(s)
}

type Circle struct {
	Side int
}

func (s *Circle) Accept(v Visitor) {
	v.VisitForCircle(s)
}

type Rectangle struct {
	Side int
}

func (s *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(s)
}

func CreateShape(shapeType ShapeType, side int) Shape {
	switch shapeType {
	case SquareType:
		return &Square{Side: side}
	case CircleType:
		return &Circle{Side: side}
	case RectangleType:
		return &Rectangle{Side: side}
	default:
		return nil
	}
}
