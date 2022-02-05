package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type info struct {
	id      int
	request string
	data    string
}

//Handles request from the info table in the database
func getInfo(db *sql.DB, request string) info {
	rows, err := db.Query("SELECT * FROM info WHERE request like '%" + request + "%'")
	handleNonFatal(err)
	var response info
	for rows.Next() {
		err = rows.Scan(&response.id, &response.request, &response.data)
		handleNonFatal(err)
	}
	err = rows.Err()
	handleNonFatal(err)
	return response
}
