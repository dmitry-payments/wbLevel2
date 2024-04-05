package main

import "fmt"

type PaymentMethod interface {
	pay(sum float64) string
}

type CreditCard struct {
	name       string
	cardNumber string
}

func (c CreditCard) pay(sum float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card (%s)", sum, c.cardNumber)
}

type PayPal struct {
	email string
}

func (p *PayPal) pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal (%s)", amount, p.email)
}

type Cryptocurrency struct {
	walletAddress string
}

func (c *Cryptocurrency) pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Cryptocurrency (%s)", amount, c.walletAddress)
}

type Item struct {
	name  string
	price float64
}

type ShoppingCart struct {
	items         []Item
	paymentMethod PaymentMethod
}

func (s *ShoppingCart) setPaymentMethod(paymentMethod PaymentMethod) {
	s.paymentMethod = paymentMethod
}

func (s *ShoppingCart) checkout() string {
	var total float64
	for _, item := range s.items {
		total += item.price
	}
	return s.paymentMethod.pay(total)
}

func main() {
	shoppingCart := &ShoppingCart{
		items: []Item{
			{"Laptop", 1500},
			{"Smartphone", 1000},
		},
	}

	creditCard := &CreditCard{"Chidozie C. Okafor", "4111-1111-1111-1111"}
	paypal := &PayPal{"chidosiky2015@gmail.com"}
	cryptocurrency := &Cryptocurrency{"0xAbcDe1234FghIjKlMnOp"}

	shoppingCart.setPaymentMethod(creditCard)
	fmt.Println(shoppingCart.checkout())

	shoppingCart.setPaymentMethod(paypal)
	fmt.Println(shoppingCart.checkout())

	shoppingCart.setPaymentMethod(cryptocurrency)
	fmt.Println(shoppingCart.checkout())
}
