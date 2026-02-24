package main

// CustomerOrderData holds the core data for a customer order.
// This struct is a plain data type with no behaviour — it is used
// by other types that each have their own responsibility.
type CustomerOrderData struct {
	OrderID  string
	ItemName string
	Amount   float64
}
