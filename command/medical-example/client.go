package command

import (
	"design-pattern-practice/command/medical-example/command"
	"design-pattern-practice/command/medical-example/invoker"
	"design-pattern-practice/command/medical-example/receiver"
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
