package mediator

import (
	"design-pattern-practice/mediator/mediator"
	"design-pattern-practice/mediator/service"
)

func main() {
	var m mediator.IMediator
	orderCreator := service.OrderCreator{Mediator: m}
	paymentCreator := service.PaymentCreator{Mediator: m}
	shippingCreator := service.ShippingCreator{Mediator: m}
	m = &mediator.OrderCreationMediator{
		OrderService:    &orderCreator,
		ShippingService: &shippingCreator,
		PaymentService:  &paymentCreator,
	}

	orderCreator.Create()
}
