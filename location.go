package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func Place() {
	config := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "Sales",
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(" ")
	}
	defer db.Close()

	var location string
	fmt.Print("Enter your city: ")
	_, err = fmt.Scanln(&location)
	if err != nil {
		log.Fatal(err)
	}

	// Querying data from the database based on user input
	var cityName string
	err = db.QueryRow("SELECT city FROM location WHERE city = ?", location).Scan(&cityName)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("City not found in database")
			os.Exit(1)
		}
		log.Fatal(err)
	}

	// If cityName is not empty, a match is found
	if cityName != "" {
		fmt.Println("City found in database:", cityName)
		// You can perform further operations here if needed
		California()
	} else {
		fmt.Println("City not found in database")
		os.Exit(1)
	}
}

/*
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
*/
