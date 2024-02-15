package main

import (
	"fmt"
	"os"
)

var location string

//var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Welcome to the WhatABurger!!...we make the fresh burgers")
	Place(location)
	Nyc()
}

func California() {
	menuList()
	var selection int
	fmt.Scanln(&selection)
	selection1 := selection
	if selection == 1 {
		fmt.Println("1. Mini bruger    1.50$")
		fmt.Println("2. Cheese burger  1.50$")
		fmt.Println("3. combos meals   4.50$")
		fmt.Fprintln(os.Stdout, []any{"*********************************************"}...)
		for selection1 == selection {
			Cart()
			break
		}
	}
	Message(selection)
}
