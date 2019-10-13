package search

import (
	"github.com/TwinProduction/discord-bot-go/cache"
	"github.com/TwinProduction/discord-bot-go/util"
	"github.com/TwinProduction/go-away"
	"github.com/bwmarrin/discordgo"
)

func Handler(bot *discordgo.Session, message *discordgo.MessageCreate, provider string, query string) bool {
	if hasKeyInCache(provider, query) {
		for _, value := range getValueFromCache(provider, query) {
			// TODO: find a way to check if the message is an error or not, and util.SendErrorMessage if it is
			bot.ChannelMessageSend(message.ChannelID, "**[cached]** "+value)
		}
		return true
	}
	if goaway.IsProfane(query) {
		util.SendErrorMessage(bot, message, "That doesn't sound like a smart thing to search...")
		putValueInCache(provider, query, []string{"That doesn't sound like a smart thing to search..."})
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

func hasKeyInCache(provider, key string) bool {
	switch provider {
	case "youtube":
		return cache.Youtube.Has(key)
	case "google":
		return cache.Google.Has(key)
	case "urban":
		return cache.Urban.Has(key)
	default:
		return false
	}
}

func getValueFromCache(provider, key string) []string {
	switch provider {
	case "youtube":
		return cache.Youtube.Get(key)
	case "google":
		return cache.Google.Get(key)
	case "urban":
		return cache.Urban.Get(key)
	default:
		return nil
	}
}

func putValueInCache(provider, key string, value []string) {
	switch provider {
	case "youtube":
		cache.Youtube.Put(key, value)
	case "google":
		cache.Google.Put(key, value)
	case "urban":
		cache.Urban.Put(key, value)
	}
}