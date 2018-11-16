package moderation

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	Constants "../../global"
	"../../util"
)

func Purge(b *discordgo.Session, m *discordgo.MessageCreate, param string)  {
	num, err := strconv.Atoi(param)
	if err != nil {
		util.SendErrorMessage(b, m, "**USAGE:** `" + Constants.COMMAND_PREFIX + "purge <number of messages>`")
		return
	}
	// TODO: Check if user is allowed to purge on the channel
	if num > 25 {
		util.SendErrorMessage(b, m, "You cannot purge more than 10 messages at once.")
		return
	}
	var messagesToPurge []string
	messages, _ := b.ChannelMessages(m.ChannelID, num, m.ID, "", "")
	for _, msg := range messages {
		messagesToPurge = append(messagesToPurge, msg.ID)
	}
	b.ChannelMessagesBulkDelete(m.ChannelID, messagesToPurge) // 1 call is better than N calls
	b.MessageReactionAdd(m.ChannelID, m.ID, Constants.EMOJI_SUCCESS)
}