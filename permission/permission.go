package permission

import (
	"../config"
	Constants "../global"
	"strings"
)


func IsAllowed(action string, userId string) bool {
	if userId == config.Config.Users.OwnerId {
		return true
	}
	switch strings.Replace(strings.ToLower(action), Constants.COMMAND_PREFIX, "", -1) {
		case "perms add": fallthrough
		case "perms remove":
			return config.Config.Users.OwnerId == userId
		default:
			if config.Config.Users.Permissions[action] == nil { // new action, add it to the list
				config.Config.Users.Permissions[action] = []string{}
				return false
			}
			return isInListOrListContainsAsterisk(config.Config.Users.Permissions[action], userId)
	}
}


func AddPermission(action string, userId string)  {
	if config.Config.Users.Permissions == nil {
		config.Config.Users.Permissions = make(map[string][]string)
	}
	if !IsAllowed(action, userId) {
		config.Config.Users.Permissions[action] = append(config.Config.Users.Permissions[action], userId)
		config.Save()
	}
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
