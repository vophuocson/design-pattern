package message

import (
	"design-pattern-practice/structural-pattern/flyweight/actor"
	"fmt"
)

type Message interface {
	DisplayMessage()
}

type TextMessage struct {
	Content string
	Actor   *actor.ActorDisplay
}

func (m *TextMessage) DisplayMessage() {
	fmt.Println("This is actor Name: " + m.Actor.Name)
	fmt.Println("This is actor Icon: " + m.Actor.Icon)
	fmt.Println("This is content Message: " + m.Content)
}

type ImageMessage struct {
	Base64 string
	Actor  *actor.ActorDisplay
}

func (m *ImageMessage) DisplayMessage() {
	fmt.Println("This is actor Name: " + m.Actor.Name)
	fmt.Println("This is actor Icon: " + m.Actor.Icon)
	fmt.Println("This is content Message: " + m.Base64)
}
