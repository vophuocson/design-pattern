package observer

import "fmt"

type Observer interface {
	Update(message string)
	GetID() string
}

type user struct {
	id string
}

func (o *user) Update(message string) {
	fmt.Println("user update by " + message)
}

func (o *user) GetID() string {
	return o.GetID()
}

func NewObserver() Observer {
	return &user{}
}
