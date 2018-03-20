package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := ""
	dbPass := ""
	dbName := ""
	dbAuth := dbUser + ":" + dbPass
	var (
		id    int
		title string
	)
	db, err := sql.Open("mysql", dbAuth+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select postID, Title from posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, title)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
