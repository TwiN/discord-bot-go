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
)


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
		switch cmd {
			case "say": say(b, m, query)
			case "shrug": b.ChannelMessageSend(m.ChannelID, m.Author.Mention()+": ¯\\_(ツ)_/¯"); b.ChannelMessageDelete(m.ChannelID, m.ID)
			case "purge": purge(b, m, query)
			case "whoami": b.ChannelMessageSend(m.ChannelID, m.Author.Username + "#" + m.Author.Discriminator)
			case "blacklist":
				if len(arguments) != 3 {
					sendErrorMessage(b, m, "**USAGE:** `" + Constants.COMMAND_PREFIX + "blacklist <add|remove> <userId>`")
					break
				}
				blacklistHandler(b, m, arguments[1], arguments[2])
			case "perms":
				if len(arguments) != 4 {
					sendErrorMessage(b, m, "**USAGE:** `" + Constants.COMMAND_PREFIX + "perms <add|remove> <cmd> <userId>`")
					break
				}
				permissionHandler(b, m, arguments[1], arguments[2], arguments[3])
			case "google": fallthrough
			case "youtube": fallthrough
			case "urban":
				searchHandler(b, m, cmd, query)
		}
	}
}


func permissionHandler(b *discordgo.Session, m *discordgo.MessageCreate, action string, cmd string, userId string) {
	switch strings.ToLower(action) {
		case "add":
			permission.AddPermission(cmd, userId)
			sendSuccessMessage(b, m, "Permissions for '" + cmd + "' has been granted to userId " + userId)
		case "remove":
			permission.RemovePermission(cmd, userId)
			sendSuccessMessage(b, m, "Permissions for '" + cmd + "' has been removed from userId " + userId)
		default:
			sendErrorMessage(b, m, "Invalid action.")
	}
}


func blacklistHandler(b *discordgo.Session, m *discordgo.MessageCreate, action string, userId string) {
	switch strings.ToLower(action) {
		case "add":
			permission.Blacklist(userId)
			sendSuccessMessage(b, m, "UserId " + userId + " has been added to the blacklist")
		case "remove":
			permission.Unblacklist(userId)
			sendSuccessMessage(b, m, "UserId " + userId + " has been removed from the blacklist")
		default:
			sendErrorMessage(b, m, "Invalid action.")
	}
}


func searchHandler(b *discordgo.Session, m *discordgo.MessageCreate, provider string, query string) {
	if cache.Has(provider, query) {
		for _, value := range cache.Get(provider, query) {
			b.ChannelMessageSend(m.ChannelID, "**[cached]** " + value)
		}
		return
	}
	if goaway.IsProfane(query) {
		b.ChannelMessageSend(m.ChannelID, "That doesn't sound like a smart thing to search...")
		cache.Put(provider, query, []string{"That doesn't sound like a smart thing to search..."})
		return
	}
	switch provider {
		case "youtube":
			search.YoutubeSearch(b, m, query)
		case "google":
			search.GoogleSearch(b, m, query)
		case "urban":
			search.UrbanDictionarySearch(b, m, query)
	}
}


func say(b *discordgo.Session, m *discordgo.MessageCreate, what string) {
	if what == "" {
		b.ChannelMessageSend(m.ChannelID, "**USAGE:** `" + Constants.COMMAND_PREFIX + "say <what>`")
		return
	}
	b.ChannelMessageSend(m.ChannelID, what)
}


func purge(b *discordgo.Session, m *discordgo.MessageCreate, param string)  {
	num, err := strconv.Atoi(param)
	if err != nil {
		sendErrorMessage(b, m, "**USAGE:** `" + Constants.COMMAND_PREFIX + "purge <number of messages>`")
		return
	}
	if num > 10 {
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