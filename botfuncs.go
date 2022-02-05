package main

import (
	"database/sql"
	"strings"

	discutil "github.com/MaxwellBanks/godiscordutil"
	_ "github.com/mattn/go-sqlite3"
)

var CommandMap = map[string]discutil.BotFunc{
	"info": infoMessage,
}

//Handles requests for basic information
func infoMessage(db *sql.DB, args []string) string {
	info := "No info data found for this request"
	if len(args) > 0 {
		response := getInfo(db, args[0])
		info = response.data
	}
	return strings.ReplaceAll(info, "\\n", "\n")
}
