package util

import (
	"github.com/bwmarrin/discordgo"
	Constants "../global"
)


func SendErrorMessage(b *discordgo.Session, m *discordgo.MessageCreate, msg string) {
	sendMessage(b, m, Constants.EMOJI_FAILURE, msg)
}


func SendSuccessMessage(b *discordgo.Session, m *discordgo.MessageCreate, msg string) {
	sendMessage(b, m, Constants.EMOJI_SUCCESS, msg)
}


func sendMessage(b *discordgo.Session, m *discordgo.MessageCreate, emojiId string, msg string) {
	b.MessageReactionAdd(m.ChannelID, m.ID, emojiId)
	b.ChannelMessageSend(m.ChannelID, msg)
}
