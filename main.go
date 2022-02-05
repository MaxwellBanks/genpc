package main

import (
	"database/sql"
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	discutil "github.com/MaxwellBanks/godiscordutil"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

type Env struct {
	db *sql.DB
}

type BotData struct {
	Flag   string
	DBPath string
}

var Token string

// Sets up discord bot token and seeds random generator
func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}

// Returns bot data for use in other functions
// I'm sure there's a better way to store constants,
// But this seems like a reasonable approach for now
func GetBotData() (BotFlag string, DBPath string) {
	BotFlag, DBPath = "!", "data.db"
	return
}

// Logs and exits in case of fatal error
func handleFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Generates bot's message to send
func (env *Env) messageCreate(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {
	botFlag, dbPath := GetBotData()
	// Ignore invalid messages
	if discutil.IsOwnMessage(s, m) || !discutil.IsCommand(m.Content, botFlag) {
		return
	}
	var BotMessage string
	command, args := discutil.ParseMessage(m.Content, botFlag)
	Message = discutil.CommandToFunc(m.Content)

}

func main() {
	_, dbPath := GetBotData()
	dg, err := discordgo.New("Bot " + Token)
	handleFatal(err)

	db, err := sql.Open("sqlite3", dbPath)
	handleFatal(err)
	defer db.Close()

	env := &Env{db: db}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(env.messageCreate)

	// Bot only cares about messages
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Create and begin listening on websocket connection
	err = dg.Open()
	handleFatal(err)

	// Wait for termination signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
