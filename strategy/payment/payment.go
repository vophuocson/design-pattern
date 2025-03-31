package payment

import "fmt"

type PayMentStrategy interface {
	Pay(amount float64)
}

type CreditCardPayment struct {
	cardNumber string
}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Thanh toán %.2f bằng thẻ tín dụng: %s\n", amount, c.cardNumber)
}

// PayPalPayment là một chiến lược thanh toán bằng PayPal
type PayPalPayment struct {
	email string
}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Thanh toán %.2f bằng PayPal: %s\n", amount, p.email)
}

func PaymentFactory(method string, credential string) PayMentStrategy {
	switch method {
	case "credit":
		return &CreditCardPayment{cardNumber: credential}
	case "paypal":
		return &PayPalPayment{email: credential}
	default:
		return nil
	}
}
