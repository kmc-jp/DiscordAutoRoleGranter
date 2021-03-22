package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Settings Setting

func init() {
	var err error

	b, err := ioutil.ReadFile(SettingsFilePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &Settings)
	if err != nil {
		panic(err)
	}

	if Settings.Discord.Token == "" {
		fmt.Println("No token provided. Please run: airhorn -t <bot token>")
		return
	}
}

func main() {
	discord, err := NewDiscord()
	if err != nil {
		panic(err)
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	discord.Session.AddHandler(discord.addRole)

	{
		var intent = discordgo.IntentsGuildMembers
		// In this example, we only care about receiving message events.
		discord.Session.Identify.Intents = intent
	}

	// Open a websocket connection to Discord and begin listening.
	err = discord.Session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc

	// Cleanly close down the Discord session.
	discord.Session.Close()
}
