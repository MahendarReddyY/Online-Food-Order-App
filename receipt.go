package main

import (
	"fmt"
	"math/rand"
	"time"
)

func receipt() {
	time.Sleep(5 * time.Second)
	fmt.Printf("Hurray !!! \n Done with Payment\n")
	fmt.Println("......thanks for ordering....!")
	fmt.Println("#########ORDER SUMMARY############")
	fmt.Println("List of items prices:", totals)
	fmt.Printf("Amount Paid: $%v\n", PaidAmount)
	fmt.Printf("Return change $%v:", change)
	fmt.Println("your order number is ", rand.Intn(100))
	fmt.Println("Promocode Applied", promo)
	fmt.Println("The final price of the item after promocode Applied:", Finalp)
	fmt.Println("updated Retrun change :", change2)
}
