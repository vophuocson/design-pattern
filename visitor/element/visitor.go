package element

type Visitor interface {
	VisitForSquare(s *Square)
	VisitForCircle(s *Circle)
	VisitForRectangle(s *Rectangle)
}
