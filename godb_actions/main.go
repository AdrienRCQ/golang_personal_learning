package main

import (
	"database/sql"
	"fmt"
)

func main() {
	var sql_password, sql_login string
	sql_address := "127.0.0.1:3306"
	fmt.Println("Golang x DataBase")
	// Setup connection for my db
	fmt.Scanln(&sql_login)
	fmt.Println("Enter your login")
	fmt.Scanln(&sql_password)
	fmt.Println("Enter your password")
	sql_credentials := sql_login + ":" + sql_password + "@tcp(" + sql_address + ")/test"

	db, err := sql.Open("mysql", sql_credentials)

	// if there is an error opening the connection, handle it :
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	defer db.Close()

}
