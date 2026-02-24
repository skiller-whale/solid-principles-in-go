package main

import (
	"fmt"
	"os"
	"strings"
)

func getInvoiceFilename(order CustomerOrderData) string {
	return "INV-" + strings.ReplaceAll(order.OrderID, "_", "-")
}

func formatOrder(order CustomerOrderData) string {
	return fmt.Sprintf("%s | %s | %.2f", order.OrderID, order.ItemName, order.Amount)
}

func logOrder(order CustomerOrderData) {
	path := "_invoices/" + getInvoiceFilename(order)
	err := os.WriteFile(path, []byte(formatOrder(order)), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error logging order: %v\n", err)
	}
}
