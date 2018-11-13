package permission

import (
	"../config"
	Constants "../global"
	"strings"
)


func IsAllowed(action string, userId string) bool {
	if userId == config.Config.Users.OwnerId || isInListOrListContainsAsterisk(config.Config.Users.Admins, userId) {
		return true
	}
	switch strings.Replace(strings.ToLower(action), Constants.COMMAND_PREFIX, "", -1) {
		case "blacklist": fallthrough
		case "unblacklist": fallthrough
		case "perms add": fallthrough
		case "perms remove": // This is handled by the if at the start of the function... TODO: remove
			return config.Config.Users.OwnerId == userId || isInListOrListContainsAsterisk(config.Config.Users.Admins, userId)
		default:
			if config.Config.Users.Permissions[action] == nil { // new action, add it to the list
				config.Config.Users.Permissions[action] = []string{}
				return false
			}
			return isInListOrListContainsAsterisk(config.Config.Users.Permissions[action], userId)
	}
}


func Blacklist(userId string) bool {
	if config.Config.Users.BlackList == nil {
		config.Config.Users.BlackList = []string{}
		return false
	}
	if userId != config.Config.Users.OwnerId && !isInListOrListContainsAsterisk(config.Config.Users.Admins, userId) && !IsBlacklisted(userId) {
		config.Config.Users.BlackList = append(config.Config.Users.BlackList, userId)
		config.Save()
		return true
	}
	return false
}


func Unblacklist(userId string) bool {
	if config.Config.Users.BlackList == nil {
		config.Config.Users.BlackList = []string{}
		return false
	}
	for i, u := range config.Config.Users.BlackList {
		if u == userId {
			config.Config.Users.BlackList = append(config.Config.Users.BlackList[:i], config.Config.Users.BlackList[i+1:]...)
			config.Save()
			return true
		}
	}
	return false
}


func IsBlacklisted(userId string) bool {
	if config.Config.Users.BlackList == nil {
		config.Config.Users.BlackList = []string{}
		return false
	}
	return isInListOrListContainsAsterisk(config.Config.Users.BlackList, userId)
}


func AddPermission(action string, userId string) bool {
	if config.Config.Users.Permissions == nil {
		config.Config.Users.Permissions = make(map[string][]string)
	}
	if !IsAllowed(action, userId) { // makes sure the user doesn't already have that permission
		config.Config.Users.Permissions[action] = append(config.Config.Users.Permissions[action], userId)
		config.Save()
		return true
	}
	return false
}


func RemovePermission(action string, userId string) bool {
	if config.Config.Users.Permissions == nil {
		return true // If it's nil, then obviously there's nothing to remove
	}
	for i, u := range config.Config.Users.Permissions[action] {
		if u == userId {
			config.Config.Users.Permissions[action] = append(config.Config.Users.Permissions[action][:i], config.Config.Users.Permissions[action][i+1:]...)
			config.Save()
			return true
		}
	}
	return false
}


func isInListOrListContainsAsterisk(haystack []string, needle string) bool {
	for _, element := range haystack {
		if element == needle || element == "*" {
			return true
		}
	}
	return false
}
