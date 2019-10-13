package moderation

import (
	"github.com/TwinProduction/discord-bot-go/global"
	"github.com/TwinProduction/discord-bot-go/util"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func Purge(bot *discordgo.Session, message *discordgo.MessageCreate, param string) {
	num, err := strconv.Atoi(param)
	if err != nil {
		util.SendErrorMessage(bot, message, "**USAGE:** `"+global.CommandPrefix+"purge <number of messages>`")
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
	bot.MessageReactionAdd(message.ChannelID, message.ID, Constants.EmojiSuccess)
}
