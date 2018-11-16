package moderation

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"../../util"
	"../../permission"
)


func BlacklistHandler(b *discordgo.Session, m *discordgo.MessageCreate, action string, userId string) {
	switch strings.ToLower(action) {
	case "add":
		if permission.Blacklist(userId) {
			util.SendSuccessMessage(b, m, "UserId " + userId + " has been added to the blacklist")
		} else {
			util.SendErrorMessage(b, m, "Couldn't add that user to the blacklist!")
		}
	case "remove":
		if permission.Unblacklist(userId) {
			util.SendSuccessMessage(b, m, "UserId " + userId + " has been removed from the blacklist")
		} else {
			util.SendErrorMessage(b, m, "There is no user with that id in the blacklist")
		}
	default:
		util.SendErrorMessage(b, m, "Invalid action.")
	}
}
