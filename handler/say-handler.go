package handler

import (
	"strings"
	"github.com/bwmarrin/discordgo"
)

func SayHandler(bot *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == bot.State.User.ID {
		return
	}
	if strings.HasPrefix(message.Content, "!say ") {
		bot.ChannelMessageSend(message.ChannelID, strings.Replace(message.Content, "!say ", "", 1))
	}
}
