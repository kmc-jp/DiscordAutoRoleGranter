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

func (d *discordHandler) removeRole(discord *discordgo.Session, event *discordgo.GuildMemberUpdate) {
	var check bool

	for _, r := range event.Roles {
		if r == Settings.Discord.RoleID {
			check = true
			break
		}
	}
	if !check {
		return
	}

	var err error
	err = discord.GuildMemberRoleRemove(
		Settings.Discord.GuildID,
		event.User.ID,
		Settings.Discord.RemoveRoleID,
	)

	if err != nil {
		fmt.Printf("Error: failed to remove Role\n%s\n", err.Error())
		return
	}

	return
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
