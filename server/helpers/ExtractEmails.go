package helpers

import (
	"regexp"
)

func ExtractEmails(htmlContent string, emailRegex *regexp.Regexp) []string {
	return emailRegex.FindAllString(htmlContent, -1)
}