package invoker

import "design-pattern-practice/command/medical-example/command"

type Button struct {
	Command command.Command
}

func (b *Button) SetCommand(c command.Command) {
	b.Command = c
}

func (b *Button) Press() {
	b.Command.Execute()
}
