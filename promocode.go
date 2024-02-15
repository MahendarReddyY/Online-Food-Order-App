package main

import (
	"fmt"
)

var promo string
var codemap []string
var code int
var finalPrice float32
var finalp int

func Promo() {
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
		//finalp = strconv.AppendFloat(total * .50, 2)
		//finalPrice = finalp
		fmt.Println("Now you can see the final price of the order")
	} else {
		fmt.Println("invalid Copuon code:")
	}
}
