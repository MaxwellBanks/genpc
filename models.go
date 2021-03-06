package main

import (
	"database/sql"
	"strconv"

	discutil "github.com/MaxwellBanks/godiscordutil"
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

type cydanger struct {
	id        int
	threshold int
	effect    string
}

func getCyDanger(db *sql.DB, args []string) (string, int, int, string) {
	charCyphers, err := strconv.Atoi(args[1])
	if err != nil {
		return "Error", -1, -1, "Error"
	}
	maxCyphers, err := strconv.Atoi(args[2])
	if err != nil {
		return "Error", -1, -1, "Error"
	}
	if charCyphers <= maxCyphers {
		return "0", 0, 60, "No effect."
	}
	roll := 0
	for i := 0; i < charCyphers-maxCyphers; i++ {
		roll += (getRandomIndex(100) + 10)
	}
	var upperThreshold int
	var lowerThreshold int
	var effect string
	err = db.QueryRow("SELECT threshold, effect FROM cydanger where threshold > "+strconv.Itoa(int(roll))+" limit 1").Scan(&upperThreshold, &effect)
	handleNonFatal(err)
	err = db.QueryRow("SELECT threshold FROM cydanger where threshold < " + strconv.Itoa(int(roll)) + " order by threshold desc limit 1").Scan(&lowerThreshold)
	handleNonFatal(err)
	return strconv.Itoa(roll), lowerThreshold + 1, upperThreshold, effect
}

type quirk struct {
	id        int
	threshold int
	effect    string
}

func getQuirk(db *sql.DB) (string, int, int, string) {
	roll := getRandomIndex(100)
	var upperThreshold int
	var lowerThreshold int
	var effect string
	err := db.QueryRow("SELECT threshold, effect FROM quirk where threshold > "+strconv.Itoa(int(roll))+" limit 1").Scan(&upperThreshold, &effect)
	handleNonFatal(err)
	err = db.QueryRow("SELECT threshold FROM quirk where threshold < " + strconv.Itoa(int(roll)) + " order by threshold desc limit 1").Scan(&lowerThreshold)
	handleNonFatal(err)
	return strconv.Itoa(roll), lowerThreshold + 1, upperThreshold, effect
}

type cypher struct {
	id       int
	name     string
	die      int
	modifier int
	methods  string
	effect   string
}

func getItemLevel(die int, mod int) int {
	return getRandomIndex(die) + mod
}

func getCypher(db *sql.DB) (string, int, string, string) {
	index := getRandomIndex(getTableSize(db, "cypher"))
	rows, err := db.Query("SELECT * FROM cypher WHERE id like '%" + strconv.Itoa(index) + "%'")
	handleNonFatal(err)
	var response cypher
	for rows.Next() {
		err = rows.Scan(&response.id, &response.name, &response.die, &response.modifier, &response.methods, &response.effect)
		handleNonFatal(err)
	}
	err = rows.Err()
	handleNonFatal(err)
	return response.name, getItemLevel(response.die, response.modifier), response.methods, response.effect
}

type artifact struct {
	id        int
	name      string
	die       int
	modifier  int
	methods   string
	effect    string
	depletion string
}

func getArtifact(db *sql.DB) (string, int, string, string, string) {
	index := getRandomIndex(getTableSize(db, "artifact"))
	rows, err := db.Query("SELECT * FROM artifact WHERE id like '%" + strconv.Itoa(index) + "%'")
	handleNonFatal(err)
	var response artifact
	for rows.Next() {
		err = rows.Scan(&response.id, &response.name, &response.die, &response.modifier, &response.methods, &response.effect, &response.depletion)
		handleNonFatal(err)
	}
	err = rows.Err()
	handleNonFatal(err)
	return response.name, getItemLevel(response.die, response.modifier), response.methods, response.effect, response.depletion
}

type threshold struct {
	id            int
	summary       string
	rollthreshold int
	effect        string
}

func getThresholds(db *sql.DB) discutil.RawTable {
	rows, err := db.Query("SELECT * FROM threshold")
	handleNonFatal(err)
	var linedata threshold
	table := [][]string{
		{"Level", "Difficulty", "Threshold", "Effect"},
	}
	for rows.Next() {
		err = rows.Scan(&linedata.id, &linedata.summary, &linedata.rollthreshold, &linedata.effect)
		handleNonFatal(err)
		line := []string{
			strconv.Itoa(linedata.id),
			linedata.summary,
			strconv.Itoa(linedata.rollthreshold),
			linedata.effect,
		}
		table = append(table, line)
	}
	err = rows.Err()
	handleNonFatal(err)
	return table
}

type kronk struct {
	id       int
	reaction string
}

func getKronk(db *sql.DB) string {
	index := getRandomIndex(getTableSize(db, "kronk"))
	rows, err := db.Query("SELECT * FROM kronk WHERE id like '%" + strconv.Itoa(index) + "%'")
	handleNonFatal(err)
	var response kronk
	for rows.Next() {
		err = rows.Scan(&response.id, &response.reaction)
		handleNonFatal(err)
	}
	err = rows.Err()
	handleNonFatal(err)
	return string(response.reaction)
}

type effect struct {
	id          int
	level       string
	effect      string
	description string
}

func getEffects(db *sql.DB) discutil.RawTable {
	rows, err := db.Query("SELECT * FROM effect")
	handleNonFatal(err)
	var linedata effect
	table := [][]string{
		{"Type", "Effect", "Description"},
	}
	for rows.Next() {
		err = rows.Scan(&linedata.id, &linedata.level, &linedata.effect, &linedata.description)
		handleNonFatal(err)
		line := []string{
			linedata.level,
			linedata.effect,
			linedata.description,
		}
		table = append(table, line)
	}
	err = rows.Err()
	handleNonFatal(err)
	return table
}

type specialroll struct {
	id          int
	roll        int
	effect      string
	description string
}

func getSpecialRolls(db *sql.DB) discutil.RawTable {
	rows, err := db.Query("SELECT * FROM special_roll")
	handleNonFatal(err)
	var linedata specialroll
	table := [][]string{
		{"Roll", "Effect", "Description"},
	}
	for rows.Next() {
		err = rows.Scan(&linedata.id, &linedata.roll, &linedata.effect, &linedata.description)
		handleNonFatal(err)
		line := []string{
			strconv.Itoa(linedata.roll),
			linedata.effect,
			linedata.description,
		}
		table = append(table, line)
	}
	err = rows.Err()
	handleNonFatal(err)
	return table
}
