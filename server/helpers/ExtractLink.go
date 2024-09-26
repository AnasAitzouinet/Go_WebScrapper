package helpers

import (
	"strings"
)

// Helper functions to extract content based on patterns
func ExtractLinks(htmlContent string) []string {
	var links []string
	doc := strings.Split(htmlContent, "<a href=")
	for _, segment := range doc[1:] {
		link := strings.SplitN(segment, `"`, 2)[0]
		if strings.HasPrefix(link, "http") {
			links = append(links, link)
		}
	}
	return links
}
