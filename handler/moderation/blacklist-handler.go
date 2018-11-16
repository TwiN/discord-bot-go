package moderation

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"../../util"
	"../../permission"
)


func BlacklistHandler(bot *discordgo.Session, message *discordgo.MessageCreate, action string, userId string) {
	switch strings.ToLower(action) {
	case "add":
		if permission.Blacklist(userId) {
			util.SendSuccessMessage(bot, message, "UserId " + userId + " has been added to the blacklist")
		} else {
			util.SendErrorMessage(bot, message, "Couldn't add that user to the blacklist!")
		}
	case "remove":
		if permission.Unblacklist(userId) {
			util.SendSuccessMessage(bot, message, "UserId " + userId + " has been removed from the blacklist")
		} else {
			util.SendErrorMessage(bot, message, "There is no user with that id in the blacklist")
		}
	default:
		util.SendErrorMessage(bot, message, "Invalid action.")
	}
}
