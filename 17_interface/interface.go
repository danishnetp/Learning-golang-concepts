package main

import (
	"fmt"
)

// Paymenter defines behavior required by any payment gateway.
type Paymenter interface {
	pay(amount float64)
}

// payment is a service that delegates charging to an injected gateway.
type payment struct {
	gateway Paymenter
}

// makePayment forwards the payment request to the active gateway.
func (p payment) makePayment(amount float64) {
	p.gateway.pay(amount)
}

// PayPal is one implementation of Paymenter.
type PayPal struct{}

// pay processes a payment through PayPal.
func (p PayPal) pay(amount float64) {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
}

// Stripe is another implementation of Paymenter.
type Stripe struct{}

// pay processes a payment through Stripe.
func (s Stripe) pay(amount float64) {
	fmt.Printf("Processing Stripe payment of $%.2f\n", amount)
}

func main() {
	// Select PayPal at runtime and execute the same payment workflow.
	paypalPayment := payment{gateway: PayPal{}}
	paypalPayment.makePayment(100.0)

	// Swap to Stripe without changing payment service logic.
	stripePayment := payment{gateway: Stripe{}}
	stripePayment.makePayment(200.0)
}
