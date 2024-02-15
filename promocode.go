package main

import (
	"fmt"
)

var promo string
var code int
var finalPrice float32
var finalp int

func Promo() {
	var codemap = []string{}
	codemap[0] = "DF345TK"
	codemap[1] = "E1QN8W"
	fmt.Println("Enter a Promo Code Here:")
	fmt.Scanln(&promo)

	for ok := range codemap {
		code = 0
		code += ok
	}
	if promo == codemap[code] {
		fmt.Println("Promocode applied sucessfully")
		fmt.Println("The final amount is:")
		finalp = (total * 50)
		finalPrice = float32(finalp / 100)
		fmt.Println("Now you can see the final price of the order", finalPrice)
	} else {
		fmt.Println("In valid Promo code ad your total cart price is :", total)
		receipt()
	}
}
