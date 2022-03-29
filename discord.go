package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type DiscordSetting struct {
	Token               string
	GuildID             string
	RoleID              string
	RemoveTriggerRoleID []string
}

type discordHandler struct {
	Session  *discordgo.Session
	settings DiscordSetting
}

func NewDiscord(settings DiscordSetting) (discordHandler, error) {
	d, err := discordgo.New("Bot " + settings.Token)
	return discordHandler{d, settings}, err
}

func (d discordHandler) Start() error {
	d.Session.AddHandler(d.addRole)
	d.Session.AddHandler(d.removeRole)

	{
		var intent = discordgo.IntentsGuildMembers
		// In this example, we only care about receiving message events.
		d.Session.Identify.Intents = intent
	}

	// Open a websocket connection to Discord and begin listening.
	err := d.Session.Open()
	return errors.Wrap(err, "OpenSession")
}

func (d discordHandler) removeRole(discord *discordgo.Session, event *discordgo.GuildMemberUpdate) {
	for _, r := range event.Roles {
		for _, s := range d.settings.RemoveTriggerRoleID {
			if r == s {
				err := discord.GuildMemberRoleRemove(
					d.settings.GuildID,
					event.User.ID,
					d.settings.RoleID,
				)

				if err != nil {
					fmt.Printf("Error: failed to remove Role\n%s\n", err.Error())
					return
				}
			}
		}
	}
}

func (d discordHandler) addRole(discord *discordgo.Session, event *discordgo.GuildMemberAdd) {
	err := discord.GuildMemberRoleAdd(
		d.settings.GuildID,
		event.User.ID,
		d.settings.RoleID,
	)

	if err != nil {
		fmt.Printf("Error: failed to Add Role\n%s\n", err.Error())
		return
	}
}
