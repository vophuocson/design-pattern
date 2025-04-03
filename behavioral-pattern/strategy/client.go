package strategy

import (
	"design-pattern-practice/behavioral-pattern/strategy/order"
	"design-pattern-practice/behavioral-pattern/strategy/payment"
)

func main() {
	payment := payment.PaymentFactory("credit", "123456789")
	order := order.Order{}
	order.SetPayment(payment)
	order.Pay()
}
