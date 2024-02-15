package main

import (
	"fmt"
	"math/rand"
	"time"
)

var total int

func receipt() {
	time.Sleep(5 * time.Second)
	fmt.Printf("Hurray !!! \n Done with Payment\n")
	fmt.Println("......thanks for ordering....!")
	fmt.Println("#########ORDER SUMMARY############")
	fmt.Println("List of items prices:", totals)
	fmt.Printf("Amount Paid: $%v\n", PaidAmount)
	fmt.Printf("Take your change $%v:", change)
	fmt.Println("your order number is ", rand.Intn(100))
}
