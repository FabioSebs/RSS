package utils

import "strings"

func ValidateTitle(title string) bool {
	title = strings.ToLower(title)
	for _, keyword := range KEYWORDS {
		if strings.Contains(title, keyword) {
			return true
		}
	}
	return false
}
