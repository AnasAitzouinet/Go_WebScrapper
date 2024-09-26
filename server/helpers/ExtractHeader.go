package helpers

import (
	"strings"
)

func ExtractHeaders(htmlContent string) []string {
	var headers []string
	for _, headerTag := range []string{"h1", "h2", "h3", "h4", "h5", "h6"} {
		doc := strings.Split(htmlContent, "<"+headerTag+">")
		for _, segment := range doc[1:] {
			header := strings.SplitN(segment, "</"+headerTag+">", 2)[0]
			headers = append(headers, header)
		}
	}
	return headers
}