package util

import (
	"github.com/TwinProduction/discord-bot-go/cache"
	"github.com/bwmarrin/discordgo"
)

func GetChannelNameById(bot *discordgo.Session, id string) string {
	if !cache.Channel.Has(id) {
		channel, _ := bot.Channel(id)
		cache.Channel.Put(id, []string{channel.Name})
		return channel.Name
	}
	return cache.Channel.Get(id)[0]
}

func GetGuildNameById(bot *discordgo.Session, id string) string {
	if !cache.Guild.Has(id) {
		guild, _ := bot.Guild(id)
		cache.Guild.Put(id, []string{guild.Name})
		return guild.Name
	}
	return cache.Guild.Get(id)[0]
}
