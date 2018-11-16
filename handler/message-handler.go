package handler

import (
	"strings"
	"github.com/bwmarrin/discordgo"
	Constants "../global"
	"../permission"
	"../util"
	"./roleplay"
	"./search"
	"./moderation"
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
		Execute:     func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
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
		Execute:     func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			b.ChannelMessageSend(m.ChannelID, m.Author.Username + "#" + m.Author.Discriminator)
			return true
		},
	},
	"pat": {
		category:    "roleplay",
		description: "Sends a GIF or an image of a character patting the head of another character",
		Execute:     func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			roleplay.Pat(b, m)
			return true
		},
	},
	"hug": {
		category:    "roleplay",
		description: "Sends a GIF or an image of a character hugging another character",
		Execute:     func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			roleplay.Hug(b, m)
			return true
		},
	},
	"greet": {
		category:    "roleplay",
		description: "Sends a GIF or an image of a character greeting another character",
		Execute:     func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			roleplay.Greet(b, m)
			return true
		},
	},
	"youtube": {
		category:    "search",
		description: "Returns the top YouTube search results for the given query",
		Execute:     func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			return search.SearchHandler(b, m, cmd, query)
		},
	},
	"google": {
		category:    "search",
		description: "Returns the top Google search results for the given query",
		Execute:     func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			return search.SearchHandler(b, m, cmd, query)
		},
	},
	"urban": {
		category:    "search",
		description: "Returns the UrbanDictionary definition of the given query",
		Execute:     func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			return search.SearchHandler(b, m, cmd, query)
		},
	},
	"purge": {
		category:    "moderation",
		description: "Removes N messages from the current channel",
		Execute:     func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			moderation.Purge(b, m, query)
			return true
		},
	},
	"blacklist": {
		category:    "moderation",
		description: "Manages blacklisted users",
		Execute:     func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			if len(arguments) != 3 {
				util.SendErrorMessage(b, m, "**USAGE:** `" + Constants.COMMAND_PREFIX + "blacklist <add|remove> <userId>`")
				return false
			}
			moderation.BlacklistHandler(b, m, arguments[1], arguments[2])
			return true
		},
	},
	"perms": {
		category:    "moderation",
		description: "Manages permissions",
		Execute:     func(b *discordgo.Session, m *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			if len(arguments) != 4 {
				util.SendErrorMessage(b, m, "**USAGE:** `" + Constants.COMMAND_PREFIX + "perms <add|remove> <cmd> <userId>`")
				return false
			}
			moderation.PermissionHandler(b, m, arguments[1], arguments[2], arguments[3])
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
		if cmd == "help" {
			help(b, m)
			return
		}
		if !permission.IsAllowed(cmd, m.Author.ID) {
			util.SendErrorMessage(b, m, "You have insufficient permissions")
			return
		}
		commandInfo, isKeyPresent := commands[cmd]
		if isKeyPresent {
			commandInfo.Execute(b, m, cmd, query, arguments)
		}
	}
}


func help(b *discordgo.Session, m *discordgo.MessageCreate) {
	output := "\n"
	for commandName, commandInfo := range commands {
		output += "__**" + Constants.COMMAND_PREFIX + commandName + "**__\n" +
			"**description:** " + commandInfo.description + "\n" +
			"**category:** " + commandInfo.category + "\n\n"
	}
	msg := &discordgo.MessageEmbed{}
	msg.Title = "List of commands available"
	msg.Description = output
	b.ChannelMessageSendEmbed(m.ChannelID, msg)
}


func swapAlias(cmd string) string {
	switch cmd {
		case "g": return "google"
		case "yt": return "youtube"
	}
	return cmd
}