package moderation

import (
	"github.com/TwinProduction/discord-bot-go/permission"
	"github.com/TwinProduction/discord-bot-go/util"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func PermissionHandler(bot *discordgo.Session, message *discordgo.MessageCreate, action string, cmd string, userId string) {
	var target string
	if userId != "*" {
		u, err := bot.User(userId)
		if err != nil {
			util.SendErrorMessage(bot, message, ":warning: There is no user with that id")
			return
		}
		target = u.Username
	} else {
		target = "all users"
	}
	switch strings.ToLower(action) {
	case "add":
		if permission.AddPermission(cmd, userId) {
			util.SendSuccessMessage(bot, message, "Permissions for '"+cmd+"' has been granted to "+target)
		} else {
			util.SendErrorMessage(bot, message, ":warning: User passed as parameter already has access to the given command.")
		}
	case "remove":
		if permission.RemovePermission(cmd, userId) {
			util.SendSuccessMessage(bot, message, "Permissions for '"+cmd+"' has been removed from "+target)
		} else {
			util.SendErrorMessage(bot, message, ":warning: User passed as parameter already doesn't have access to the given command.")
		}
	default:
		util.SendErrorMessage(bot, message, ":warning: Invalid action.")
	}
}
