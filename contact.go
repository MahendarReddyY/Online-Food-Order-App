package main

import (
	"fmt"
)

func Message(selection int) {
	for {
		if selection == 2 {
			fmt.Println("Buy one get one Burger ..! $3min purchase required ")
			fmt.Println("Get a iced coffee with purchase of any breakfast meals")
		}
		if selection == 3 {
			fmt.Println("Phone : +1 403-907-0076")
			fmt.Println("Email: whatAburger@customer.care")
		}

		if selection > 3 {
			fmt.Println("There is no option", selection, "choose an option from 1 to 3")
			Place(location)
		}
		California()
	}
}
