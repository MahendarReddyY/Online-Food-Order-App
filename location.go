package main

import (
	"fmt"
	"strings"
)

var spot = map[string]string{
	"1": "Dallas",
	"2": "Newyork",
	"3": "Washington Dc",
	"4": "NewJersey",
}
var location string

func Place() {
	fmt.Println("Our locations are Dallas, Newyork, Washington Dc, NewJersey")
	fmt.Println("Enter the Location:")
	fmt.Scanln(&location)
	location = strings.ToLower(strings.TrimSpace(location))

	var found bool
	for _, value := range spot {
		if strings.ToLower(value) == location {
			fmt.Printf("Thanks for choosing  %s \n", value)
			found = true
			California()
		}
	}
	if !found {
		fmt.Println("Location not found.")
	}
}
