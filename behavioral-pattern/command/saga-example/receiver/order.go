package receiver

import "fmt"

type CreationOrder struct {
}

func (o *CreationOrder) CreateOrder() error {
	fmt.Println("this is CreateOrder")
	return nil
}

func (o *CreationOrder) CompensateOrder() {
	fmt.Println("this is CompensateOrder")
}

func (o *CreationOrder) VerifyConsumer() error {
	fmt.Println("this is VerifyConsumer()")
	return nil
}

func (o *CreationOrder) CreateTicket() error {
	fmt.Println("this is CreateTicket ticket")
	return nil
}

func (o *CreationOrder) CompensateTicket() error {
	fmt.Println("this is CompensateTicket ticket")
	return nil
}

func (c *CreationOrder) AuthenticateCard() error {
	fmt.Println("this is AuthenticateCard card")
	return nil
}

func (c *CreationOrder) Verify() error {
	fmt.Println("this is Verify")
	return nil
}

func (c *CreationOrder) ApproveTicket() error {
	fmt.Println("this is ApproveTicket")
	return nil
}

func (c *CreationOrder) ApproveOrder() error {
	fmt.Println("this is ApproveOrder")
	return nil
}
