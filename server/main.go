package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/chromedp/chromedp"

	"scrapper/server/helpers"
	"scrapper/server/models"
)

// scrapeWithChromedp is a function that scrapes the data from the webpage using chromedp
func scrapeWithChromedp(url string, body models.RequestBody) (models.ResponseBody, error) {
	// Create a new ChromeDP context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Ensure the context is cleaned up after completion
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Initialize variables to hold scraped data
	var htmlContent, pageTitle string

	// Run ChromeDP actions to scrape the webpage content
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.Title(&pageTitle),
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		return models.ResponseBody{}, err
	}

	responseBody := models.ResponseBody{}
	responseBody.Title = pageTitle

	// Now process the HTML content to find links, paragraphs, images, etc.
	emailRegex := regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`)

	// If 'Link' or 'All' is set, scrape links
	if body.Link || body.All {
		responseBody.Links = helpers.ExtractLinks(htmlContent)
	}

	// If 'Paragraph' or 'All' is set, scrape paragraphs
	if body.Paragraph || body.All {
		responseBody.Paragraphs = helpers.ExtractParagraphs(htmlContent)
	}

	if body.Image || body.All {
		responseBody.Images = helpers.ExtractImages(htmlContent, body)
	}

	if body.Header || body.All {
		responseBody.Headers = helpers.ExtractHeaders(htmlContent)
	}

	if body.Email || body.All {
		responseBody.Emails = helpers.ExtractEmails(htmlContent, emailRegex)
	}


	// If 'All' is set, scrape images, headers, and emails
	if body.All {
		responseBody.Images = helpers.ExtractImages(htmlContent, body)
		responseBody.Headers = helpers.ExtractHeaders(htmlContent)
		responseBody.Emails = helpers.ExtractEmails(htmlContent, emailRegex)
	}

	return responseBody, nil
}





// handleRequest is the handler for the / endpoint
func handleRequest(w http.ResponseWriter, r *http.Request) {
	body := models.RequestBody{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set body.Link and body.Paragraph to true if body.All is set
	if body.All {
		body.Link = true
		body.Paragraph = true
	}

	// Scrape the data using chromedp
	responseBody, err := scrapeWithChromedp(body.Url, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the scraped data as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}

// main is the entry point for the application
func main() {
	// Start the server on port 8080
	fmt.Println("Starting server on port http://localhost:8080")

	// Handle the / endpoint
	http.HandleFunc("/", handleRequest)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
