package main

func processOrder(orderID, item string, amount float64, cardType string) {
	order := CustomerOrderData{OrderID: orderID, ItemName: item, Amount: amount}
	logOrder(order)
	calculator := InvoiceCalculator{Amount: order.Amount, CardType: cardType}
	printer := InvoicePrinter{Order: order, Calculator: calculator}
	printer.Print()
}

func main() {
	processOrder("abc_defghi_jkl", "Lunch", 38.20, "amex")
	processOrder("mno_pqrstu_vwx", "Dinner", 24.99, "visa")
}
