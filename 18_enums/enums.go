package main

import "fmt"

// OrderStatus is an enum-like custom type for order lifecycle states.
// Go does not have native enums, so a typed int + const + iota is the common pattern.
type OrderStatus int

const (
	// OrderPending is the initial status when an order is created.
	OrderPending OrderStatus = iota
	// OrderShipped means the order has left the warehouse.
	OrderShipped
	// OrderDelivered means the order reached the customer.
	OrderDelivered
	// OrderCancelled means the order will not be fulfilled.
	OrderCancelled
)

// changeOrderStatus demonstrates passing strongly typed enum-like values.
func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	// Example transitions in an order workflow.
	changeOrderStatus(OrderPending)
	changeOrderStatus(OrderShipped)
}
