package main

import (
	"fmt"
	"strings"
)

func Place(location string) string {
	for {
		fmt.Println("Enter the Location or Area Code")
		fmt.Scanln(&location)
		loc := strings.ToUpper(location)
		switch loc {
		case "Newyork", "940":
			fmt.Printf("Thanks for choosing this %v location\n", loc)
			California()
		case "California", "920":
			fmt.Printf("Thanks for choosing this %v location", loc)
			Nyc()
		default:
			fmt.Printf("Sorry! our store is not in your %v location.\n", loc)
			main()
		}
		return loc
	}
}
