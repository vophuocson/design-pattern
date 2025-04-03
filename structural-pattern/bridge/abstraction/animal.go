package abstraction

import "design-pattern-practice/structural-pattern/bridge/implementation"

type Animal interface {
	HowMove()
}

type Person struct {
	MoveLogic implementation.MoveLogic
}

func (a *Person) HowMove() {
	a.MoveLogic.Move()
}

type Fish struct {
	MoveLogic implementation.MoveLogic
}

func (a *Fish) HowMove() {
	a.MoveLogic.Move()
}

type Bird struct {
	MoveLogic implementation.MoveLogic
}

func (a *Bird) HowMove() {
	a.MoveLogic.Move()
}
