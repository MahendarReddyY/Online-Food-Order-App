package main

import (
	"fmt"
)

var PaidAmount int
var change int

func payment() {
	fmt.Println("Welcome to the Payment Gate way")
	fmt.Println("Amount paid:")
	fmt.Scanln(&PaidAmount)
	total := sum(totals)
	fmt.Println("Total Price is:", total)
	// Calculate change
	Promo()
	change = PaidAmount - total
	fmt.Println("Take your change:", change)

	go receipt()

}

func sum(totals []int) int {
	total := 0
	for _, num := range totals {
		total += num
	}
	return total
}
