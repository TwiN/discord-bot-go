package util

import (
	"github.com/bwmarrin/discordgo"
	Constants "../global"
)


func SendErrorMessage(bot *discordgo.Session, message *discordgo.MessageCreate, msg string) {
	sendMessage(bot, message, Constants.EMOJI_FAILURE, msg)
}


func SendSuccessMessage(bot *discordgo.Session, message *discordgo.MessageCreate, msg string) {
	sendMessage(bot, message, Constants.EMOJI_SUCCESS, msg)
}


func sendMessage(bot *discordgo.Session, message *discordgo.MessageCreate, emojiId string, msg string) {
	bot.MessageReactionAdd(message.ChannelID, message.ID, emojiId)
	bot.ChannelMessageSend(message.ChannelID, msg)
}
