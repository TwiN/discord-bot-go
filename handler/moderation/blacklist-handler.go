package moderation

import (
	"fmt"
	"github.com/TwinProduction/discord-bot-go/permission"
	"github.com/TwinProduction/discord-bot-go/util"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func BlacklistHandler(bot *discordgo.Session, message *discordgo.MessageCreate, action string, userId string) {
	var target string
	if userId != "*" {
		u, err := bot.User(userId)
		if err != nil {
			util.SendErrorMessage(bot, message, ":warning: There is no user with that id")
			return
		}
		target = u.Username
	} else {
		target = "All users"
	}
	switch strings.ToLower(action) {
	case "add":
		if permission.Blacklist(userId) {
			util.SendSuccessMessage(bot, message, fmt.Sprintf("%s has been added to the blacklist", target))
		} else {
			util.SendErrorMessage(bot, message, ":warning: Couldn't add that user to the blacklist!")
		}
	case "remove":
		if permission.Unblacklist(userId) {
			util.SendSuccessMessage(bot, message, fmt.Sprintf("%s has been removed from the blacklist", target))
		} else {
			util.SendErrorMessage(bot, message, ":warning: There is no user with that id in the blacklist")
		}
	default:
		util.SendErrorMessage(bot, message, ":warning: Invalid action.")
	}
}
