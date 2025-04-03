package command

import "design-pattern-practice/behavioral-pattern/command/medical-example/receiver"

type Command interface {
	Execute()
}

type OnCommand struct {
	Receiver receiver.Device
}

func (c *OnCommand) Execute() {
	c.Receiver.On()
}

type OffCommand struct {
	Receiver receiver.Device
}

func (c *OffCommand) Execute() {
	c.Receiver.On()
}
