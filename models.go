package main

import (
	"database/sql"
	"strconv"

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

type oddity struct {
	id          int
	description string
}

func getOddity(db *sql.DB) string {
	index := getRandomIndex(getTableSize(db, "oddity"))
	rows, err := db.Query("SELECT * FROM oddity WHERE id like '%" + strconv.Itoa(index) + "%'")
	handleNonFatal(err)
	var response oddity
	for rows.Next() {
		err = rows.Scan(&response.id, &response.description)
		handleNonFatal(err)
	}
	err = rows.Err()
	handleNonFatal(err)
	return string(response.description + ".")
}
