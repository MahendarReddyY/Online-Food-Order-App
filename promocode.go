package main

import (
	"fmt"
)

var promo string

var code []string

// var finalPrice float32
var Finalp int
var codemap = map[string]string{"1": "QTLY23", "2": "OFFER50", "3": "NEW25", "4": "Flat25"}
var change2 float32

//var promoverif = make([]string, 1)

func Promo() {
	fmt.Println("Enter a Promo Code Here:")
	fmt.Scanln(&promo)
	if len(codemap) > 0 {
		for _, ok := range codemap {
			code = append(code, ok)
		}

		for _, ok1 := range code {

			if promo != ok1 {
				fmt.Println(" In valid Promocode ")
				fmt.Println("Your total cart price is remains same :", Total)
			} else {
				fmt.Println("Promocode applied sucessfully")
				Finalp = (Total * 50) / 100
				//finalPrice = float32(finalp / 100)
				change2 = float32(PaidAmount) - float32(Finalp)
				fmt.Printf("Now you can see the final price : %v \n and remaining balance of order: %v \n ", Finalp, change2)
			}
		}
	}
	//code = append(code, ok)
}
