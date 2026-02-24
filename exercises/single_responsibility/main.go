package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type CustomerOrder struct {
	OrderID  string
	ItemName string
	Amount   float64
}

var cardFeeRates = map[string]float64{
	"visa":       0.015,
	"mastercard": 0.01,
	"amex":       0.03,
}

const invoiceWidth = 40
const salesTaxRate = 0.2

func (o CustomerOrder) invoiceNumber() string {
	return "INV-" + strings.ReplaceAll(o.OrderID, "_", "-")
}

func (o CustomerOrder) formatOrder() string {
	return fmt.Sprintf("%s | %.2f", o.ItemName, o.Amount)
}

func (o CustomerOrder) formatValue(key string, value float64) string {
	return fmt.Sprintf("%-*s%12.2f", invoiceWidth-12, key+":", value)
}

func (o CustomerOrder) SalesTax() float64 {
	return salesTaxRate * o.Amount
}

func (o CustomerOrder) CardFees(cardType string) float64 {
	return cardFeeRates[cardType] * o.Amount
}

func (o CustomerOrder) TotalCost(cardType string) float64 {
	return o.Amount + o.SalesTax() + o.CardFees(cardType)
}

func (o CustomerOrder) LogOrder() {
	path := "_invoices/" + o.invoiceNumber()
	err := os.WriteFile(path, []byte(o.formatOrder()), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error logging order: %v\n", err)
	}
}

func (o CustomerOrder) PrintInvoice(cardType string) {
	bar := strings.Repeat("-", invoiceWidth)
	now := time.Now().Format("15:04 02-Jan-2006")

	center := func(s string) string {
		pad := (invoiceWidth - len(s)) / 2
		return strings.Repeat(" ", pad) + s
	}
	rjust := func(s string) string {
		return fmt.Sprintf("%*s", invoiceWidth, s)
	}

	fmt.Println()
	fmt.Println(center("INVOICE"))
	fmt.Println(bar)
	fmt.Println(rjust(now))
	fmt.Println(rjust("No: " + o.invoiceNumber()))
	fmt.Println(bar)
	fmt.Println(o.formatOrder())
	fmt.Println(bar)
	fmt.Println(o.formatValue("Sub Total", o.Amount))
	fmt.Println(o.formatValue(fmt.Sprintf("Sales Tax (%v%%)", salesTaxRate), o.SalesTax()))
	fmt.Println(o.formatValue("Card Fees", o.CardFees(cardType)))
	fmt.Println(o.formatValue("Total", o.TotalCost(cardType)))
	fmt.Println(bar)
}

func processOrder(orderID, item string, amount float64, card string) {
	order := CustomerOrder{OrderID: orderID, ItemName: item, Amount: amount}
	order.LogOrder()
	order.PrintInvoice(card)
}

func main() {
	processOrder("abc_defghi_jkl", "Lunch", 38.20, "amex")
	processOrder("mno_pqrstu_vwx", "Dinner", 24.99, "visa")
}
