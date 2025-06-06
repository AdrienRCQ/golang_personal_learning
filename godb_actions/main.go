package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
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
	sql_credentials := sql_login + ":" + sql_password + "@tcp(" + sql_address + ")/openbar"

	db, err := sql.Open("mysql", sql_credentials)

	// if there is an error opening the connection, handle it :
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connexion réussie à la base de données MySQL!")

	// Exemple : lecture de données
	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		log.Fatalf("Erreur lors de la requête : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var nom_tache string
		var date string
		err := rows.Scan(&nom_tache, &date)
		if err != nil {
			log.Fatalf("Erreur de lecture de ligne : %v", err)
		}
		fmt.Printf("Name: %s, Date: %s\n", nom_tache, date)
	}

	// defer the close till after the main function has finished
	defer db.Close()

}
