package search

import (
	"github.com/TwinProduction/discord-bot-go/cache"
	"github.com/TwinProduction/discord-bot-go/global"
	"github.com/TwinProduction/discord-bot-go/scraper/google"
	"github.com/bwmarrin/discordgo"
)

func GoogleSearch(bot *discordgo.Session, message *discordgo.MessageCreate, query string) {
	const Command = global.CommandPrefix + "google"

	if len(query) == 0 {
		bot.ChannelMessageSend(message.ChannelID, "**USAGE:** `"+Command+" <search terms>`")
	} else {
		bot.UpdateStatus(1, "| :mag_right: '"+query+"' on Google")
		results := google.Scrape(query)
		cache.Google.Put(query, results)
		for _, url := range results {
			bot.ChannelMessageSend(message.ChannelID, url)
		}
		bot.UpdateStatus(0, "")
	}
}
