package util

import (
	"github.com/bwmarrin/discordgo"
	"../cache"
)


func GetChannelNameById(bot *discordgo.Session, id string) string {
	if !cache.Has("channel", id) {
		channel, _ := bot.Channel(id)
		cache.Put("channel", id, []string{channel.Name})
		return channel.Name
	}
	return cache.Get("channel", id)[0]
}


func GetGuildNameById(bot *discordgo.Session, id string) string {
	if !cache.Has("guild", id) {
		guild, _ := bot.Guild(id)
		cache.Put("guild", id, []string{guild.Name})
		return guild.Name
	}
	return cache.Get("guild", id)[0]
}
