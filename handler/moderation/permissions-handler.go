package moderation

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"../../util"
	"../../permission"
)

func PermissionHandler(b *discordgo.Session, m *discordgo.MessageCreate, action string, cmd string, userId string) {
	switch strings.ToLower(action) {
		case "add":
			if permission.AddPermission(cmd, userId) {
				util.SendSuccessMessage(b, m, "Permissions for '" + cmd + "' has been granted to userId " + userId)
			} else {
				util.SendErrorMessage(b, m, "User passed as parameter already has access to the given command.")
			}
		case "remove":
			if permission.RemovePermission(cmd, userId) {
				util.SendSuccessMessage(b, m, "Permissions for '" + cmd + "' has been removed from userId " + userId)
			} else {
				util.SendErrorMessage(b, m, "User passed as parameter already doesn't have access to the given command.")
			}
		default:
			util.SendErrorMessage(b, m, "Invalid action.")
	}
}