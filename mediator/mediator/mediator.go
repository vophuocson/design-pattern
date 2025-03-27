package mediator

type Service interface {
	Execute() error
}

type IMediator interface {
	Execute() error
}

type OrderCreationMediator struct {
	OrderService    Service
	PaymentService  Service
	ShippingService Service
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
