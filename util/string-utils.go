package util

import "strings"


func PadRight(s string, expectedLength int, pad string) string {
	for {
		s += pad
		if len(s) > expectedLength {
			return s[0:expectedLength]
		}
	}
}


func MentionToUserId(s string) string {
	s = strings.Trim(s, " ")
	if strings.HasPrefix(s, "<@!") && strings.HasSuffix(s, ">") {
		return s[3:len(s)-1]
	} else if strings.HasPrefix(s, "<@") && strings.HasSuffix(s, ">") {
		return s[2:len(s)-1]
	}
	return s
}
