package bridge

import (
	"design-pattern-practice/structural-pattern/bridge/abstraction"
	"design-pattern-practice/structural-pattern/bridge/implementation"
)

func Main() {
	var moveLogic implementation.MoveLogic
	walk := implementation.Walk{}
	moveLogic = &walk

	var animal abstraction.Animal
	person := abstraction.Person{MoveLogic: moveLogic}
	animal = &person
	animal.HowMove()

	Fly := implementation.Fly{}
	moveLogic = &Fly

	bird := abstraction.Bird{MoveLogic: moveLogic}
	animal = &bird
	animal.HowMove()
}
