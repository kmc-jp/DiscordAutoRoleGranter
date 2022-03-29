package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	var settings = DiscordSetting{
		Token:   os.Getenv("DISCORD_TOKEN"),
		GuildID: os.Getenv("DISCORD_GUILD_ID"),
		RoleID:  os.Getenv("DISCORD_ROLE_ID"),
	}

	var triggerRoles = []string{}

	for _, r := range strings.Split(os.Getenv("DISCORD_TRIGGER_ROLES"), ",") {
		triggerRoles = append(triggerRoles, strings.TrimSpace(r))
	}

	settings.RemoveTriggerRoleID = triggerRoles

	discord, err := NewDiscord(settings)
	if err != nil {
		panic(err)
	}

	go func() {
		err = discord.Start()
		if err != nil {
			panic(err)
		}
	}()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc

	// Cleanly close down the Discord session.
	discord.Session.Close()
}
