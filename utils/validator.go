package utils

import "strings"

func ValidateTitle(title string) bool {
	title = strings.ToLower(title)
	for _, keyword := range KEYWORDS_EN {
		if strings.Contains(title, keyword) {
			return true
		}
	}
	return false
}

func ValidateTitleID(title string) bool {
	title = strings.ToLower(title)
	for _, keyword := range KEYWORDS_ID {
		if strings.Contains(title, keyword) {
			return true
		}
	}
	return false
}
