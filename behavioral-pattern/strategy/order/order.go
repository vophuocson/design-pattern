package order

import "design-pattern-practice/behavioral-pattern/strategy/payment"

type Order struct {
	payment payment.PayMentStrategy
}

func (o *Order) SetPayment(payment payment.PayMentStrategy) {
	o.payment = payment
}

func (o *Order) Pay() {
	o.payment.Pay(20.000)
}
