package main

var cardFeeRates = map[string]float64{
	"visa":       0.015,
	"mastercard": 0.01,
	"amex":       0.03,
}

const salesTaxRate = 0.2

type InvoiceCalculator struct {
	Amount   float64
	CardType string
}

func (c InvoiceCalculator) Total() float64 {
	return c.Amount + c.SalesTax() + c.CardFees()
}

func (c InvoiceCalculator) SalesTax() float64 {
	return salesTaxRate * c.Amount
}

func (c InvoiceCalculator) CardFees() float64 {
	return cardFeeRates[c.CardType] * c.Amount
}

func (c InvoiceCalculator) TaxRate() float64 {
	return salesTaxRate
}
