package main

import (
	"fmt"
)

var PaidAmount int
var change int
var Total int

func payment() {
	fmt.Println("Welcome to the Payment Gate way")
	fmt.Println("Amount paid:")
	fmt.Scanln(&PaidAmount)
	Total = sum(totals)
	fmt.Println("Total Price is:", Total)

	// Calculate change
	change = PaidAmount - Total
	fmt.Println("Take your change:", change)

	Promo()

	receipt()

}

func sum(totals []int) int {
	total := 0
	for _, num := range totals {
		total += num
	}
	return total
}
