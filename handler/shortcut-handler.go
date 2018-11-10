package handler

import (
	"strings"
	"github.com/bwmarrin/discordgo"
	Constants "../global"
)

func ShortcutConverterHandler(bot *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == bot.State.User.ID {
		return
	}
	if strings.HasPrefix(message.Content, Constants.COMMAND_PREFIX) {
		var command = strings.Split(message.Content, " ")[0]
		var params = strings.Trim(strings.Replace(message.Content, command, "", 1), " ")
		var realCommand = ""
		switch strings.Replace(command, Constants.COMMAND_PREFIX, "", 1) {
			case "yt": realCommand = "youtube"; break
			case "g": realCommand = "google"; break
			default: return
		}
		message.Content = Constants.COMMAND_PREFIX + realCommand + " " + params
	}
}
