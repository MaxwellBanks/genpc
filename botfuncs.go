package main

import (
	"database/sql"
	"math/rand"
	"strconv"
	"strings"

	discutil "github.com/MaxwellBanks/godiscordutil"
	_ "github.com/mattn/go-sqlite3"
)

type GenFunc func(*sql.DB) string

var CommandMap = map[string]discutil.BotFunc{
	"info": infoMessage,
	"gen":  genMessage,
}

// Gets random index for tables
func getRandomIndex(max int) int {
	return rand.Intn(max) + 1
}

// Check size of table for random generation checks
func getTableSize(db *sql.DB, tablename string) int {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM " + tablename).Scan(&count)
	handleNonFatal(err)
	return count
}

// Handles requests for basic information
func infoMessage(db *sql.DB, args []string) string {
	response := "No info data found for this request"
	if len(args) > 0 {
		info := getInfo(db, args[0])
		response = info.data
	}
	return strings.ReplaceAll(response, "\\n", "\n")
}

// Checks that a generation request has enough arguments
func verifyGenRequest(args []string) bool {
	switch args[0] {
	case "cydanger":
		if len(args) < 2 {
			return false
		}
	default:
		if len(args) == 0 {
			return false
		}
	}
	return true
}

// Handles requests for random generation of data
func genMessage(db *sql.DB, args []string) string {
	response := "Unable to generate this request"
	if !verifyGenRequest(args) {
		return response
	}
	// I'd like to make this generic, but there are some
	// Odd idiosyncracies with the typing of the map
	// Value (it's typed as GenFunc despite its value
	// Returning a string), so until I figure it out
	// Switch statements it is
	//
	// var genMap = map[string]GenFunc{
	// 	"oddity": getOddity,
	// }
	// info := genMap[args[0]]
	// return strings.ReplaceAll(info, "\\n", "\n")

	switch args[0] {
	case "oddity":
		response = getOddity(db)
	case "cydanger":
		roll, lowerBound, upperBound, effect := getCyDanger(db, args)
		if lowerBound > 200 {
			response = "Roll value of " + roll + " was above " + strconv.Itoa(lowerBound) +
				"\n\n" + "Effect: " + effect
		} else {
			response = "Roll value of " + roll + " was between " + strconv.Itoa(lowerBound) +
				" and " + strconv.Itoa(upperBound) + "\n\n" + "Effect: " + effect
		}
	}
	return strings.ReplaceAll(response, "\\n", "\n")

}
