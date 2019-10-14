package search

import (
	"github.com/TwinProduction/discord-bot-go/cache"
	"github.com/TwinProduction/discord-bot-go/global"
	"github.com/TwinProduction/discord-bot-go/scraper/urban"
	"github.com/bwmarrin/discordgo"
)

func UrbanDictionarySearch(bot *discordgo.Session, message *discordgo.MessageCreate, query string) {
	const Command = global.CommandPrefix + "urban"

	if len(query) == 0 {
		bot.ChannelMessageSend(message.ChannelID, "**USAGE:** `"+Command+" <search terms>`")
	} else {
		bot.UpdateStatus(1, "| :mag_right: '"+query+"' on UrbanDictionary")
		result := "**Urban Dictionary search result for `" + query + "`:**" + urban.Scrape(query)
		cache.Urban.Put(query, []string{result})
		bot.ChannelMessageSend(message.ChannelID, result)
		bot.UpdateStatus(0, "")
	}
}
