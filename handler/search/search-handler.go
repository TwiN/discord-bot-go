package search

import (
	"github.com/TwinProduction/discord-bot-go/cache"
	"github.com/TwinProduction/discord-bot-go/util"
	"github.com/TwinProduction/go-away"
	"github.com/bwmarrin/discordgo"
)

func SearchHandler(bot *discordgo.Session, message *discordgo.MessageCreate, provider string, query string) bool {
	if cache.Has(provider, query) {
		for _, value := range cache.Get(provider, query) {
			// TODO: find a way to check if the message is an error or not, and util.SendErrorMessage if it is
			bot.ChannelMessageSend(message.ChannelID, "**[cached]** "+value)
		}
		return true
	}
	if goaway.IsProfane(query) {
		util.SendErrorMessage(bot, message, "That doesn't sound like a smart thing to search...")
		cache.Put(provider, query, []string{"That doesn't sound like a smart thing to search..."})
		return true
	}
	switch provider {
	case "youtube":
		YoutubeSearch(bot, message, query)
	case "google":
		GoogleSearch(bot, message, query)
	case "urban":
		UrbanDictionarySearch(bot, message, query)
	}
	return true
}
