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
