package command

import (
	"design-pattern-practice/command/command"
	"design-pattern-practice/command/invoker"
	"design-pattern-practice/command/receiver"
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
