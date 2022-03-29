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
	for _, r := range event.Roles {
		for _, s := range Settings.Discord.RemoveTriggerRoleID {
			if r == s {
				err := discord.GuildMemberRoleRemove(
					Settings.Discord.GuildID,
					event.User.ID,
					Settings.Discord.RoleID,
				)

				if err != nil {
					fmt.Printf("Error: failed to remove Role\n%s\n", err.Error())
					return
				}
			}
		}
	}
}

func (d *discordHandler) addRole(discord *discordgo.Session, event *discordgo.GuildMemberAdd) {
	err := discord.GuildMemberRoleAdd(
		Settings.Discord.GuildID,
		event.User.ID,
		Settings.Discord.RoleID,
	)

	if err != nil {
		fmt.Printf("Error: failed to Add Role\n%s\n", err.Error())
		return
	}
}
