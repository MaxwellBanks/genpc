package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	discutil "github.com/MaxwellBanks/godiscordutil"
	_ "github.com/mattn/go-sqlite3"
)

// GenFunc is a type signature for generator functions
type GenFunc func(*sql.DB) string

// CommandMap contains a map of user commands to functions
var CommandMap = map[string]discutil.BotFunc{
	"info":             infoMessage,
	"gen":              genMessage,
	"generate":         genMessage,
	"cs":               csMessage,
	"cheatsheet":       csMessage,
	"mcs":              csMobile,
	"mobilecheatsheet": csMobile,
	"kronk":            kronkMessage,
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
	case "oddity", "odditys", "oddities":
		response = getOddity(db)
	case "cydanger", "cypherdanger", "cypherdangers", "cydangers":
		roll, lowerBound, upperBound, effect := getCyDanger(db, args)
		if lowerBound > 200 {
			response = "Roll value of " + roll + " was above " + strconv.Itoa(lowerBound) +
				"\n\n" + "Effect: " + effect
			response = fmt.Sprintf(
				"Roll value of %s was above %d\n\nEffect: %s",
				roll,
				lowerBound,
				effect,
			)
		} else {
			response = fmt.Sprintf(
				"Roll value of %s was between %d and %d\n\nEffect: %s",
				roll,
				lowerBound,
				upperBound,
				effect,
			)
		}
	case "quirk", "quirks":
		roll, lowerBound, upperBound, effect := getQuirk(db)
		response = fmt.Sprintf(
			"Roll value of %s was between %d and %d\n\nEffect: %s",
			roll,
			lowerBound,
			upperBound,
			effect,
		)
	case "cypher", "cyphers":
		name, level, methods, effect := getCypher(db)
		response = fmt.Sprintf(
			"Generated Cypher\n**Level %d %s**\n%s\n\nEffect: %s",
			level,
			name,
			methods,
			effect,
		)
	case "artifact", "artifacts":
		name, level, methods, effect, depletion := getArtifact(db)
		response = fmt.Sprintf(
			"Generated Artifact\n**Level %d %s**\n%s\n\nEffect: %s\nDepletion: %s",
			level,
			name,
			methods,
			effect,
			depletion,
		)
	}
	return strings.ReplaceAll(response, "\\n", "\n")
}

// Handles cheatsheet info, will expand this to give more nuanced info in discord proper
func csMessage(db *sql.DB, args []string) string {
	response := "No data found."
	if len(args) == 0 {
		return "Cheatsheet link: https://www.thealexandrian.net/creations/numenera/numenera-cheat-sheet-final.pdf"
	}
	switch args[0] {
	case "thresholds", "threshold":
		response = fmt.Sprintf(
			"Roll Thresholds\n```\n%s\n```", discutil.GenTable(getThresholds(db)),
		)

	}
	return response
}

// Handles cheatsheet info for mobile devices
func csMobile(db *sql.DB, args []string) string {
	response := "No data found."
	if len(args) == 0 {
		return "Cheatsheet link: https://www.thealexandrian.net/creations/numenera/numenera-cheat-sheet-final.pdf"
	}
	switch args[0] {
	case "thresholds", "threshold":
		table := getThresholds(db)
		var mobileTable [][]string
		for i := range table {
			mobileTable = append(mobileTable, table[i][:len(table[i])-1])
		}
		response = fmt.Sprintf(
			"Roll Thresholds\n```\n%s\n```", discutil.GenTable(mobileTable),
		)

	}
	return response
}

func kronkMessage(db *sql.DB, args []string) string {
	// reaction := getKronk(db)
	// response := fmt.Sprintf("Kronk's Reaction: %s", reaction)
	// return response
	return getKronk(db)
}
