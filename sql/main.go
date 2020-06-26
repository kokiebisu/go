package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// This essentially means that, every time you 
// query your database, you are using a connection
// from a pool of connections that have been set up
// on application startup. These connections are reused,
// time and time again, and this subsequently means you 
// aren’t creating and destroying a new connection every
// time you perform a query.

type Tag struct {
	ID int `json:"id"`
	description string `json:"name"`
}

func main() {
	fmt.Println("this is the mysql tutorial")

	// db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	db, err := sql.Open("mysql", "ken:kronos111@tcp(127.0.0.1:3306)/go")
	if err != nil {
		panic(err)
	}

	// defer the close until the main function is finished
	defer db.Close()

	// INSERT
	// insert, err := db.Query("insert into test values (1, 'first description')")

	// if err != nil {
	// 	panic(err)
	// }

	// defer insert.Close()

	results, err := db.Query("select id, description from test");

	if err != nil {
		panic(err.Error()) // this is the proper way of handling Errors
	}

	// Next() scans each row
	for results.Next() {
		var tag Tag // create a new object to put the row data in
		err := results.Scan(&tag.ID, &tag.description) // must match
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("tag description:" + tag.description)

	}
	var differentTag Tag
	newError := db.QueryRow("select * from test where id=?", 1).Scan(&differentTag.ID, &differentTag.description)
	if newError != nil {
			panic(err.Error())
	}
	fmt.Println("single row description" + differentTag.description)





}