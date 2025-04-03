package sagaexample

import (
	"design-pattern-practice/behavioral-pattern/command/saga-example/command"
	"design-pattern-practice/behavioral-pattern/command/saga-example/orchestrator"
	"design-pattern-practice/behavioral-pattern/command/saga-example/receiver"
	"fmt"
)

func main() {
	receiver := receiver.CreationOrder{}
	saga := orchestrator.Orchestrator{}
	err := saga.Execute(&command.OrderCreationCommand{Receiver: &receiver})
	if err != nil {
		fmt.Println("Execute command error")
	}
	saga.Execute(&command.ConsumerVerificationCommand{Receiver: &receiver})
	if err != nil {
		fmt.Println("Execute command error")
	}
	saga.Execute(&command.TicketCreationCommand{Receiver: &receiver})
	if err != nil {
		fmt.Println("Execute command error")
	}
	saga.Execute(&command.CardAuthenticationCommand{Receiver: &receiver})
	if err != nil {
		fmt.Println("Execute command error")
	}
	saga.Execute(&command.OrderVerificationCommand{Receiver: &receiver})
	if err != nil {
		fmt.Println("Execute command error")
	}
	saga.Execute(&command.TicketApprovalCommand{Receiver: &receiver})
	if err != nil {
		fmt.Println("Execute command error")
	}
	saga.Execute(&command.OrderApprovalCommand{Receiver: &receiver})
	if err != nil {
		fmt.Println("Execute command error")
	}
}
