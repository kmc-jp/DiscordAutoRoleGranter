package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type discordHandler struct {
	Session *discordgo.Session
}

func NewDiscord() (discordHandler, error) {
	d, err := discordgo.New("Bot " + Settings.Discord.Token)
	return discordHandler{d}, err
}

func (d *discordHandler) addRole(discord *discordgo.Session, event *discordgo.GuildMemberAdd) {
	var err error

	err = discord.GuildMemberRoleAdd(
		Settings.Discord.GuildID,
		event.User.ID,
		Settings.Discord.RoleID,
	)

	if err != nil {
		fmt.Printf("Error: failed to Add Role\n%s\n", err.Error())
		return
	}

	return
}
