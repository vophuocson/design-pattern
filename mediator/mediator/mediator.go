package mediator

type service interface {
	Execute() error
}

type IMediator interface {
	Execute() error
}

type OrderCreationMediator struct {
	OrderService    service
	PaymentService  service
	ShippingService service
}

func (m *OrderCreationMediator) Execute() error {
	err := m.PaymentService.Execute()
	if err != nil {
		return err
	}
	err = m.OrderService.Execute()
	if err != nil {
		return err
	}
	err = m.ShippingService.Execute()
	if err != nil {
		return err
	}
	return nil
}
