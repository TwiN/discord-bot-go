package handler

import (
	"strings"
	"strconv"
	"github.com/bwmarrin/discordgo"
	"github.com/TwinProduction/go-away"
	"./search"
	Constants "../global"
	"../cache"
	"../permission"
	"./roleplay"
)

type CommandInfo struct {
	category      string
	description   string
	Execute       func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool
}



var commands = map[string]CommandInfo {
	"shrug": {
		category:    "misc",
		description: "¯\\_(ツ)_/¯",
		Execute: func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			_, e := b.ChannelMessageSend(m.ChannelID, m.Author.Mention()+": ¯\\_(ツ)_/¯")
			if e != nil {
				return false
			}
			return b.ChannelMessageDelete(m.ChannelID, m.ID) == nil
		},
	},
	"whoami": {
		category:    "misc",
		description: "Replies with the username of the user followed by the discriminator, e.g. `Twin#9089`.",
		Execute: func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			b.ChannelMessageSend(m.ChannelID, m.Author.Username + "#" + m.Author.Discriminator)
			return true
		},
	},
	"pat": {
		category:    "roleplay",
		description: "¯\\_(ツ)_/¯",
		Execute: func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			roleplay.Pat(b, m)
			return true
		},
	},
	"hug": {
		category:    "roleplay",
		description: "¯\\_(ツ)_/¯",
		Execute: func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			roleplay.Hug(b, m)
			return true
		},
	},
	"greet": {
		category:    "roleplay",
		description: "¯\\_(ツ)_/¯",
		Execute: func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			roleplay.Greet(b, m)
			return true
		},
	},
	"youtube": {
		category:    "search",
		description: "¯\\_(ツ)_/¯",
		Execute: func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			return searchHandler(b, m, cmd, query)
		},
	},
	"google": {
		category:    "search",
		description: "¯\\_(ツ)_/¯",
		Execute: func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			return searchHandler(b, m, cmd, query)
		},
	},
	"urban": {
		category:    "search",
		description: "¯\\_(ツ)_/¯",
		Execute: func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			return searchHandler(b, m, cmd, query)
		},
	},
	"purge": {
		category:    "moderation",
		description: "Removes N messages from the current channel",
		Execute: func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			// TODO: Check if user is allowed to purge first
			purge(b, m, query)
			return true
		},
	},
	"blacklist": {
		category:    "moderation",
		description: "Manages blacklisted users",
		Execute: func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			if len(arguments) != 3 {
				sendErrorMessage(b, m, "**USAGE:** `" + Constants.COMMAND_PREFIX + "blacklist <add|remove> <userId>`")
				return false
			}
			blacklistHandler(b, m, arguments[1], arguments[2])
			return true
		},
	},
	"perms": {
		category:    "moderation",
		description: "Manages permissions",
		Execute: func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			if len(arguments) != 4 {
				sendErrorMessage(b, m, "**USAGE:** `" + Constants.COMMAND_PREFIX + "perms <add|remove> <cmd> <userId>`")
				return false
			}
			permissionHandler(b, m, arguments[1], arguments[2], arguments[3])
			return true
		},
	},
}


func MessageHandler(b *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == b.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, Constants.COMMAND_PREFIX) {
		arguments := strings.Split(strings.Trim(strings.Replace(m.Content, Constants.COMMAND_PREFIX, "", 1), " "), " ")
		query := strings.Replace(m.Content, Constants.COMMAND_PREFIX + arguments[0] + " ", "", 1)
		cmd := swapAlias(strings.ToLower(arguments[0]))

		if permission.IsBlacklisted(m.Author.ID) {
			return
		}
		if !permission.IsAllowed(cmd, m.Author.ID) {
			sendErrorMessage(b, m, "You have insufficient permissions")
			return
		}
		commandInfo, isKeyPresent := commands[cmd]
		if isKeyPresent {
			commandInfo.Execute(b, m, cmd, query, arguments)
		}
	}
}


func permissionHandler(b *discordgo.Session, m *discordgo.MessageCreate, action string, cmd string, userId string) {
	switch strings.ToLower(action) {
		case "add":
			if permission.AddPermission(cmd, userId) {
				sendSuccessMessage(b, m, "Permissions for '" + cmd + "' has been granted to userId " + userId)
			} else {
				sendErrorMessage(b, m, "User passed as parameter already has access to the given command.")
			}
		case "remove":
			if permission.RemovePermission(cmd, userId) {
				sendSuccessMessage(b, m, "Permissions for '" + cmd + "' has been removed from userId " + userId)
			} else {
				sendErrorMessage(b, m, "User passed as parameter already doesn't have access to the given command.")
			}
		default:
			sendErrorMessage(b, m, "Invalid action.")
	}
}


func blacklistHandler(b *discordgo.Session, m *discordgo.MessageCreate, action string, userId string) {
	switch strings.ToLower(action) {
		case "add":
			if permission.Blacklist(userId) {
				sendSuccessMessage(b, m, "UserId " + userId + " has been added to the blacklist")
			} else {
				sendErrorMessage(b, m, "Couldn't add that user to the blacklist!")
			}
		case "remove":
			if permission.Unblacklist(userId) {
				sendSuccessMessage(b, m, "UserId " + userId + " has been removed from the blacklist")
			} else {
				sendErrorMessage(b, m, "There is no user with that id in the blacklist")
			}
		default:
			sendErrorMessage(b, m, "Invalid action.")
	}
}


func searchHandler(b *discordgo.Session, m *discordgo.MessageCreate, provider string, query string) bool {
	if cache.Has(provider, query) {
		for _, value := range cache.Get(provider, query) {
			b.ChannelMessageSend(m.ChannelID, "**[cached]** " + value)
		}
		return true
	}
	if goaway.IsProfane(query) {
		b.ChannelMessageSend(m.ChannelID, "That doesn't sound like a smart thing to search...")
		cache.Put(provider, query, []string{"That doesn't sound like a smart thing to search..."})
		return true
	}
	switch provider {
		case "youtube":
			search.YoutubeSearch(b, m, query)
		case "google":
			search.GoogleSearch(b, m, query)
		case "urban":
			search.UrbanDictionarySearch(b, m, query)
	}
	return true
}


func purge(b *discordgo.Session, m *discordgo.MessageCreate, param string)  {
	num, err := strconv.Atoi(param)
	if err != nil {
		sendErrorMessage(b, m, "**USAGE:** `" + Constants.COMMAND_PREFIX + "purge <number of messages>`")
		return
	}
	if num > 25 {
		sendErrorMessage(b, m, "You cannot purge more than 10 messages at once.")
		return
	}
	var messagesToPurge []string
	messages, _ := b.ChannelMessages(m.ChannelID, num, m.ID, "", "")
	for _, msg := range messages {
		messagesToPurge = append(messagesToPurge, msg.ID)
	}
	b.ChannelMessagesBulkDelete(m.ChannelID, messagesToPurge) // 1 call is better than N calls
	b.MessageReactionAdd(m.ChannelID, m.ID, Constants.EMOJI_SUCCESS)
}


func sendErrorMessage(b *discordgo.Session, m *discordgo.MessageCreate, msg string) {
	sendMessage(b, m, Constants.EMOJI_FAILURE, msg)
}


func sendSuccessMessage(b *discordgo.Session, m *discordgo.MessageCreate, msg string) {
	sendMessage(b, m, Constants.EMOJI_SUCCESS, msg)
}


func sendMessage(b *discordgo.Session, m *discordgo.MessageCreate, emojiId string, msg string) {
	b.MessageReactionAdd(m.ChannelID, m.ID, emojiId)
	b.ChannelMessageSend(m.ChannelID, msg)
}


func swapAlias(cmd string) string {
	switch cmd {
		case "g": return "google"
		case "yt": return "youtube"
	}
	return cmd
}