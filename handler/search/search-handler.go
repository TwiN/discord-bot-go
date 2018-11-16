package search

import (
	"github.com/bwmarrin/discordgo"
	"github.com/TwinProduction/go-away"
	"../../cache"
	"../../util"
)


func SearchHandler(b *discordgo.Session, m *discordgo.MessageCreate, provider string, query string) bool {
	if cache.Has(provider, query) {
		for _, value := range cache.Get(provider, query) {
			// TODO: find a way to check if the message is an error or not, and util.SendErrorMessage if it is
			b.ChannelMessageSend(m.ChannelID, "**[cached]** " + value)
		}
		return true
	}
	if goaway.IsProfane(query) {
		util.SendErrorMessage(b, m, "That doesn't sound like a smart thing to search...")
		cache.Put(provider, query, []string{"That doesn't sound like a smart thing to search..."})
		return true
	}
	switch provider {
	case "youtube":
		YoutubeSearch(b, m, query)
	case "google":
		GoogleSearch(b, m, query)
	case "urban":
		UrbanDictionarySearch(b, m, query)
	}
	return true
}
