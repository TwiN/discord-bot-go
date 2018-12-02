package misc

import (
	"github.com/bwmarrin/discordgo"
	"../../util"
)


func AvatarHandler(bot *discordgo.Session, message *discordgo.MessageCreate, query string) bool {
	userId := util.MentionToUserId(query)
	var avatar = message.Author.AvatarURL("512")
	if len(userId) != 0 {
		user, err := bot.User(userId)
		if err != nil {
			util.SendErrorMessage(bot, message, "**ERROR:** Invalid username")
			return false
		}
		avatar = user.AvatarURL("512")
	}
	bot.ChannelMessageSend(message.ChannelID, avatar)
	return true
}
