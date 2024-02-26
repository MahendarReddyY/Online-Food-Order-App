package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func Data() {
	config := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "Sales",
	}
	Db, Err := sql.Open("mysql", config.FormatDSN())
	if Err != nil {
		fmt.Println(Err)
	} else {
		fmt.Println(".....!!!........")
	}
	defer Db.Close()
}
