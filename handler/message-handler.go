package handler

import (
	"strings"
	"strconv"
	"github.com/bwmarrin/discordgo"
	"github.com/TwinProduction/go-away"
	"./search"
	Constants "../global"
	"../cache"
)


func MessageHandler(b *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == b.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, Constants.COMMAND_PREFIX) {
		arguments := strings.Split(strings.Trim(strings.Replace(m.Content, Constants.COMMAND_PREFIX, "", 1), " "), " ")
		cmd := arguments[0]
		query := strings.Replace(m.Content, Constants.COMMAND_PREFIX + cmd + " ", "", 1)

		println("cmd="+cmd+"; query="+query)
		switch strings.ToLower(cmd) {
			case "say": say(b, m, query)
			case "shrug": b.ChannelMessageSend(m.ChannelID, "¯\\_(ツ)_/¯")
			case "purge": purge(b, m, query)
			case "whoami": b.ChannelMessageSend(m.ChannelID, m.Author.Username + "#" + m.Author.Discriminator)

			case "google": fallthrough
			case "g": cmd = "google"; fallthrough
			case "youtube": fallthrough
			case "yt": cmd = "youtube"; fallthrough
			case "urban":
				searchHandler(b, m, cmd, query)
		}
	}
}


func searchHandler(b *discordgo.Session, m *discordgo.MessageCreate, provider string, query string) {
	if cache.Has(provider, query) {
		for _, value := range cache.Get(provider, query) {
			b.ChannelMessageSend(m.ChannelID, "**[cached]** " + value)
		}
		return
	}
	if goaway.IsProfane(query) {
		b.ChannelMessageSend(m.ChannelID, "That doesn't sound like a smart thing to search...")
		cache.Put(provider, query, []string{"That doesn't sound like a smart thing to search..."})
		return
	}

	switch strings.ToLower(provider) {
		case "youtube":
			search.YoutubeSearch(b, m, query)
		case "google":
			search.GoogleSearch(b, m, query)
		case "urban":
			search.UrbanDictionarySearch(b, m, query)
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
