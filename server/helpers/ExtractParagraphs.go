package helpers

import (
	"strings"
)

func ExtractParagraphs(htmlContent string) []string {
	var paragraphs []string
	doc := strings.Split(htmlContent, "<p>")
	for _, segment := range doc[1:] {
		paragraph := strings.SplitN(segment, "</p>", 2)[0]
		paragraphs = append(paragraphs, paragraph)
	}
	return paragraphs
}