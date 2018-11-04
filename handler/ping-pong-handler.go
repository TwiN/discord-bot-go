package handler

import (
	"github.com/bwmarrin/discordgo"
)

func PingPongHandler(bot *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == bot.State.User.ID {
		return
	}
	if message.Content == "ping" {
		bot.ChannelMessageSend(message.ChannelID, "pong")
	}
}
