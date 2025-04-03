package service

import (
	"design-pattern-practice/behavioral-pattern/mediator/mediator"
	"fmt"
)

type OrderCreator struct {
	Mediator mediator.IMediator
}

func (s *OrderCreator) Create() error {
	return s.Mediator.Execute()
}

func (s *OrderCreator) Execute() error {
	fmt.Println("This is create order")
	return nil
}

type PaymentCreator struct {
	Mediator mediator.IMediator
}

func (s *PaymentCreator) Execute() error {
	fmt.Println("This is create payment")
	return nil
}

type ShippingCreator struct {
	Mediator mediator.IMediator
}

func (s *ShippingCreator) Execute() error {
	fmt.Println("This is create CreateShipping")
	return nil
}
