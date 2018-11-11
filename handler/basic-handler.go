package handler

import (
	"strings"
	"strconv"
	"github.com/bwmarrin/discordgo"
	Constants "../global"
)


func BasicHandler(b *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == b.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, Constants.COMMAND_PREFIX) {
		arguments := strings.Split(strings.Trim(strings.Replace(m.Content, Constants.COMMAND_PREFIX, "", 1), " "), " ")
		cmd := arguments[0]
		query := strings.Replace(m.Content, Constants.COMMAND_PREFIX + cmd + " ", "", 1)

		switch strings.ToLower(cmd) {
			case "say": say(b, m, query); break
			case "shrug": b.ChannelMessageSend(m.ChannelID, "¯\\_(ツ)_/¯"); break
			case "purge": purge(b, m, query); break
			case "whoami": b.ChannelMessageSend(m.ChannelID, m.Author.Username + "#" + m.Author.Discriminator); break
		}
	}
}


func say(b *discordgo.Session, m *discordgo.MessageCreate, what string) {
	if what == "" {
		b.ChannelMessageSend(m.ChannelID, "**USAGE:** `" + Constants.COMMAND_PREFIX + "say <what>`")
		return
	}
	b.ChannelMessageSend(m.ChannelID, what)
}


func purge(b *discordgo.Session, m *discordgo.MessageCreate, param string)  {
	num, err := strconv.Atoi(param)
	if err != nil {
		sendErrorMessage(b, m, "**USAGE:** `" + Constants.COMMAND_PREFIX + "purge <number of messages>`")
		return
	}
	if num > 10 {
		sendErrorMessage(b, m, "You cannot purge more than 10 messages at once.")
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


func sendErrorMessage(b *discordgo.Session, m *discordgo.MessageCreate, errorMessage string) {
	b.MessageReactionAdd(m.ChannelID, m.ID, Constants.EMOJI_FAILURE)
	b.ChannelMessageSend(m.ChannelID, errorMessage)
}
