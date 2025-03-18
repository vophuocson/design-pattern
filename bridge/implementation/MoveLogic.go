package implementation

import "fmt"

type MoveLogic interface {
	Move()
}

type Walk struct{}

func (m *Walk) Move() {
	fmt.Println("this is the action of Walking")
}

type Fly struct{}

func (m *Fly) Move() {
	fmt.Println("this is the action of Flying")
}

type Swim struct{}

func (m *Swim) Move() {
	fmt.Println("this is the action of Swimming")
}
