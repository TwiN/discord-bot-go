package util

import (
	"github.com/TwinProduction/discord-bot-go/global"
	"github.com/bwmarrin/discordgo"
)

func SendErrorMessage(bot *discordgo.Session, message *discordgo.MessageCreate, msg string) {
	sendMessage(bot, message, global.EmojiFailure, msg)
}

func SendSuccessMessage(bot *discordgo.Session, message *discordgo.MessageCreate, msg string) {
	sendMessage(bot, message, global.EmojiSuccess, msg)
}

func sendMessage(bot *discordgo.Session, message *discordgo.MessageCreate, emojiId string, msg string) {
	bot.MessageReactionAdd(message.ChannelID, message.ID, emojiId)
	bot.ChannelMessageSend(message.ChannelID, msg)
}
