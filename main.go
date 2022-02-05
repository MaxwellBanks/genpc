package main

import (
	"flag"
	"math/rand"
	"time"

	discutil "github.com/MaxwellBanks/godiscordutil"
	"github.com/bwmarrin/discordgo"
)

var Token string
var Flag string = "!"

type botfunc func([]string) string

var CommandMap = map[string]botfunc{}

//Sets up discord bot token and seeds random generator
func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}

//Generates bot's message to send
func (env *Env) messageCreate(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {
	if discutil.IsOwnMessage(m) || !discutil.IsCommand(m.Content, Flag) {
		return
	}
	var Message string

}
