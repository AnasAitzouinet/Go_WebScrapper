package helpers

import (
	"log"
	"strings"

	"scrapper/server/models"

	"github.com/PuerkitoBio/goquery"
)

func ExtractImages(htmlContent string, body models.RequestBody) []string {
	var images []string

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		image, _ := s.Attr("src")
		images = append(images, image)
	})
	return images
}
