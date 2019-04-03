package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jmoiron/sqlx"

	// SQLite3 driver for the database
	_ "github.com/mattn/go-sqlite3"

	"log"
	"os"
	"os/signal"
	"syscall"
	"wumpus-birthday/pkg"
	"wumpus-birthday/pkg/globals"
	"wumpus-birthday/pkg/storage"
)

var token string
var dbFilePath string

func init() {
	flag.StringVar(&token, "t", "", "The Bot Token")
	flag.StringVar(&dbFilePath, "db", "./data.sqlite", "The database path")
	flag.Parse()

	if db, err := sqlx.Connect("sqlite3", dbFilePath); err != nil {
		log.Fatal("failed to open the sqlite database:", err)
	} else {
		globals.DB = db
		db.MustExec(storage.Schema)
	}
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	pkg.Register(dg)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	fmt.Printf(
		"Invite: https://discordapp.com/oauth2/authorize?client_id=%s&scope=bot\n", dg.State.User.ID)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	if err := dg.Close(); err != nil {
		log.Fatal(err)
	}
}
