package main

import (
	"fmt"
	"strings"
	"time"
)

const printWidth = 40

type InvoicePrinter struct {
	Order      CustomerOrderData
	Calculator InvoiceCalculator
}

func (p InvoicePrinter) invoiceNumber() string {
	return "INV-" + strings.ReplaceAll(p.Order.OrderID, "_", "-")
}

func (p InvoicePrinter) formatValue(key string, value float64) string {
	return fmt.Sprintf("%-*s%12.2f", printWidth-12, key+":", value)
}

func (p InvoicePrinter) Print() {
	bar := strings.Repeat("-", printWidth)
	now := time.Now().Format("15:04 02-Jan-2006")

	center := func(s string) string {
		pad := (printWidth - len(s)) / 2
		return strings.Repeat(" ", pad) + s
	}
	rjust := func(s string) string {
		return fmt.Sprintf("%*s", printWidth, s)
	}

	fmt.Println()
	fmt.Println(center("INVOICE"))
	fmt.Println(bar)
	fmt.Println(rjust(now))
	fmt.Println(rjust("No: " + p.invoiceNumber()))
	fmt.Println(bar)
	fmt.Printf("%s | %.2f\n", p.Order.ItemName, p.Order.Amount)
	fmt.Println(bar)
	fmt.Println(p.formatValue("Sub Total", p.Calculator.Amount))
	fmt.Println(p.formatValue(fmt.Sprintf("Sales Tax (%.0f%%)", p.Calculator.TaxRate()*100), p.Calculator.SalesTax()))
	fmt.Println(p.formatValue("Card Fees", p.Calculator.CardFees()))
	fmt.Println(p.formatValue("Total", p.Calculator.Total()))
	fmt.Println(bar)
}
