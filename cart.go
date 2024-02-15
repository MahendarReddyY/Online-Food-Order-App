package main

import "fmt"

var cartChoice int
var cartSelection []int
var pricemap = map[int]int{1: 2, 2: 3, 3: 4}
var totals = []int{}

func Cart() {
	for {
		fmt.Println("select your food Item into cart:")
		fmt.Println("Select any other key apart 123 to goto payment section")
		fmt.Scanln(&cartChoice)
		if cartChoice <= 3 {
			switch cartChoice {
			case 1, 2, 3:
				// Check if the selected item is in the pricemap
				if price, ok := pricemap[cartChoice]; ok {
					// If the item is in the map, add its price to totals
					totals = append(totals, price)
					fmt.Printf("Item: %v, Price: %d added to cart.\n", cartChoice, price)
					fmt.Println("Cart total:", totals)
				} else {
					fmt.Println("Invalid item selected.")
				}
			}
			continue
		}
		if cartChoice > 3 {
			payment()
		}
	}
}
