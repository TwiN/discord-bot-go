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
	Execute       func(bot *discordgo.Session, message *discordgo.MessageCreate, cmd string, query string, arguments []string) bool
}


var commands = map[string]CommandInfo {
	"shrug": {
		category:    "misc",
		description: "¯\\_(ツ)_/¯",
		Execute:     func(bot *discordgo.Session, message *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			_, e := bot.ChannelMessageSend(message.ChannelID, message.Author.Mention()+": ¯\\_(ツ)_/¯")
			if e != nil {
				return false
			}
			return bot.ChannelMessageDelete(message.ChannelID, message.ID) == nil
		},
	},
	"whoami": {
		category:    "misc",
		description: "Replies with the username of the user followed by the discriminator, e.g. `Twin#9089`.",
		Execute:     func(bot *discordgo.Session, message *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			bot.ChannelMessageSend(message.ChannelID, message.Author.Username + "#" + message.Author.Discriminator)
			return true
		},
	},
	"pat": {
		category:    "roleplay",
		description: "Sends a GIF or an image of a character patting the head of another character",
		Execute:     func(bot *discordgo.Session, message *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			roleplay.Pat(bot, message)
			return true
		},
	},
	"hug": {
		category:    "roleplay",
		description: "Sends a GIF or an image of a character hugging another character",
		Execute:     func(bot *discordgo.Session, message *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			roleplay.Hug(bot, message)
			return true
		},
	},
	"greet": {
		category:    "roleplay",
		description: "Sends a GIF or an image of a character greeting another character",
		Execute:     func(bot *discordgo.Session, message *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			roleplay.Greet(bot, message)
			return true
		},
	},
	"youtube": {
		category:    "search",
		description: "Returns the top YouTube search results for the given query",
		Execute:     func(bot *discordgo.Session, message *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			return search.SearchHandler(bot, message, cmd, query)
		},
	},
	"google": {
		category:    "search",
		description: "Returns the top Google search results for the given query",
		Execute:     func(bot *discordgo.Session, message *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			return search.SearchHandler(bot, message, cmd, query)
		},
	},
	"urban": {
		category:    "search",
		description: "Returns the UrbanDictionary definition of the given query",
		Execute:     func(bot *discordgo.Session, message *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			return search.SearchHandler(bot, message, cmd, query)
		},
	},
	"purge": {
		category:    "moderation",
		description: "Removes N messages from the current channel",
		Execute:     func(bot *discordgo.Session, message *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			moderation.Purge(bot, message, query)
			return true
		},
	},
	"blacklist": {
		category:    "moderation",
		description: "Manages blacklisted users",
		Execute:     func(bot *discordgo.Session, message *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			if len(arguments) != 3 {
				util.SendErrorMessage(bot, message, "**USAGE:** `" + Constants.COMMAND_PREFIX + "blacklist <add|remove> <userId>`")
				return false
			}
			moderation.BlacklistHandler(bot, message, arguments[1], arguments[2])
			return true
		},
	},
	"perms": {
		category:    "moderation",
		description: "Manages permissions",
		Execute:     func(bot *discordgo.Session, message *discordgo.MessageCreate, cmd string, query string, arguments []string) bool {
			if len(arguments) != 4 {
				util.SendErrorMessage(bot, message, "**USAGE:** `" + Constants.COMMAND_PREFIX + "perms <add|remove> <cmd> <userId>`")
				return false
			}
			moderation.PermissionHandler(bot, message, arguments[1], arguments[2], arguments[3])
			return true
		},
	},
}


func MessageHandler(bot *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == bot.State.User.ID {
		return
	}
	if strings.HasPrefix(message.Content, Constants.COMMAND_PREFIX) {
		arguments := strings.Split(strings.Trim(strings.Replace(message.Content, Constants.COMMAND_PREFIX, "", 1), " "), " ")
		query := strings.Replace(message.Content, Constants.COMMAND_PREFIX + arguments[0] + " ", "", 1)
		cmd := swapAlias(strings.ToLower(arguments[0]))

		if permission.IsBlacklisted(message.Author.ID) {
			return
		}
		if cmd == "help" {
			help(bot, message)
			return
		}
		if !permission.IsAllowed(cmd, message.Author.ID) {
			util.SendErrorMessage(bot, message, "You have insufficient permissions")
			return
		}
		commandInfo, isKeyPresent := commands[cmd]
		if isKeyPresent {
			commandInfo.Execute(bot, message, cmd, query, arguments)
		}
	}
}


func help(bot *discordgo.Session, message *discordgo.MessageCreate) {
	output := "\n"
	for commandName, commandInfo := range commands {
		output += "__**" + Constants.COMMAND_PREFIX + commandName + "**__\n" +
			"**description:** " + commandInfo.description + "\n" +
			"**category:** " + commandInfo.category + "\n\n"
	}
	msg := &discordgo.MessageEmbed{}
	msg.Title = "List of commands available"
	msg.Description = output
	bot.ChannelMessageSendEmbed(message.ChannelID, msg)
}


func swapAlias(cmd string) string {
	switch cmd {
		case "g": return "google"
		case "yt": return "youtube"
	}
	return cmd
}
