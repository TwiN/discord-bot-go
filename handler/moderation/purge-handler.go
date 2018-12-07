package moderation

import (
	"strconv"
	"github.com/bwmarrin/discordgo"
	Constants "../../global"
	"../../util"
)


func Purge(bot *discordgo.Session, message *discordgo.MessageCreate, param string)  {
	num, err := strconv.Atoi(param)
	if err != nil {
		util.SendErrorMessage(bot, message, "**USAGE:** `" + Constants.COMMAND_PREFIX + "purge <number of messages>`")
		return
	}
	// TODO: Check if user is allowed to purge on the channel
	if num > 25 {
		util.SendErrorMessage(bot, message, "You cannot purge more than 10 messages at once.")
		return
	}
	var messagesToPurge []string
	messages, _ := bot.ChannelMessages(message.ChannelID, num, message.ID, "", "")
	for _, msg := range messages {
		messagesToPurge = append(messagesToPurge, msg.ID)
	}
	bot.ChannelMessagesBulkDelete(message.ChannelID, messagesToPurge) // 1 call is better than N calls
	bot.MessageReactionAdd(message.ChannelID, message.ID, Constants.EMOJI_SUCCESS)
}
