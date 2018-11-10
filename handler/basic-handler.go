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
			case "purge": purge(b, m, query); break
		}

		/*if len(query) == 0 {
			bot.ChannelMessageSend(message.ChannelID, "**USAGE:** `" + COMMAND + " <search terms>`")
		} else {

		}*/

	}
}


func say(b *discordgo.Session, m *discordgo.MessageCreate, what string)  {
	if what == "" {
		b.ChannelMessageSend(m.ChannelID, "**USAGE:** `" + Constants.COMMAND_PREFIX + "say <what>`")
		return
	}
	b.ChannelMessageSend(m.ChannelID, what)
}


func purge(b *discordgo.Session, m *discordgo.MessageCreate, param string)  {
	num, err := strconv.Atoi(param)
	if err != nil {
		b.ChannelMessageSend(m.ChannelID, "**USAGE:** `" + Constants.COMMAND_PREFIX + "purge <number of messages>`")
		return
	}
	b.ChannelMessageSend(m.ChannelID, "Purging " + param + " messages")
	b.ChannelMessageDelete(m.ChannelID, m.ID)

	messages, _ := b.ChannelMessages(m.ChannelID, num, "", "", "")
	for msg := range messages {
		println(msg)
	}
	b.ChannelMessagesBulkDelete(m.ChannelID, []string{})
}
