// Write your answer here, and then test your code.

package main

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {
	var total float64 = 0
	for _, item := range cart {
		total += (item.price * float64(item.quantity))
	}
	return total
}
