package validation

import "regexp"

func IsValidUserName(name string) bool {
	if len(name) == 0 || len(name) > 32 || len(name) < 3 {
		return false
	}

	matched, _ := regexp.MatchString("^[a-zA-Z]+(?:[\\s'-][a-zA-Z]+)*$", name)

	return matched
}

func IsValidMessageContent(content string) bool {
	if len(content) == 0 || len(content) > 256 {
		return false
	}

	return true
}
