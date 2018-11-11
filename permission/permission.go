package permission

import (
	"../config"
	Constants "../global"
	"strings"
)


func IsAllowed(action string, userId string) bool {
	switch strings.Replace(strings.ToLower(action), Constants.COMMAND_PREFIX, "", -1) {
		case "purge":
			return config.Config.Users.OwnerId == userId
		default:
			return true
	}
}


func Allow() {

}