package main

import (
	"fmt"
)

var firstName string
var lastName string
var email string
var Address int

func input() {

	fmt.Println("Customer information:")
	fmt.Println("Enter your firstName:")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your lastName:")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scanln(&email)

	fmt.Println("Enter number your address:")
	fmt.Scanln(&Address)

	receipt()

}

/*
	isvalidName, isvalidEmail, isvalidTickets := Evaluated(firstName, lastName, email)

	if isvalidEmail && isvalidName && isvalidTickets {
		payment()
	} else {
		validcheck(isvalidEmail, isvalidName, isvalidTickets)
	}
}

func validcheck(isvalidEmail bool, isvalidName bool, isvalidTickets bool) {
	if !isvalidEmail {
		fmt.Println("your email address must contains @ and .com signs ")
	}
	if !isvalidName {
		fmt.Println("your name is too short ")
	}
	if !isvalidTickets {
		fmt.Println("enter required tickets values from 1 to 50 ")
	}
}

func Evaluated(firstName string, lastName string, email string) (bool, bool, bool) {
	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	return isvalidName, isvalidEmail, isvalidName
}
*/
