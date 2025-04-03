package command

import (
	"design-pattern-practice/behavioral-pattern/command/medical-example/command"
	"design-pattern-practice/behavioral-pattern/command/medical-example/invoker"
	"design-pattern-practice/behavioral-pattern/command/medical-example/receiver"
)

func main() {
	var d receiver.Device
	d = &receiver.Radio{}
	var c command.Command
	c = &command.OnCommand{Receiver: d}
	button := invoker.Button{}
	button.SetCommand(c)
	button.Press()
}
