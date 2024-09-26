
# Scrapper Project

## Description
This is a simple web scraping service built in Go, using the `chromedp` package to scrape dynamic web content and extract useful information like links, paragraphs, headers, images, and emails from a given URL. The server exposes an API that accepts a request body with options on what to scrape and returns the extracted data in JSON format.

## Project Structure

```bash
.
├── server                # The main directory for the server-side code
│   ├── helpers           # Helper functions for processing scraped data
│   │   └── ExtractImages.go  # Contains helper functions for extracting images
│   │   └── ExtractHeaders.go  # Contains helper functions for extracting headers
│   │   └── ExtractEmails.go  # Contains helper functions for extracting emails
│   │   └── ExtractLinks.go  # Contains helper functions for extracting links
│   │   └── ExtractParagraphs.go  # Contains helper functions for extracting paragraphs
│   ├── models            # Data models for requests and responses
│   │   └── models.go     # Contains RequestBody and ResponseBody structures
│   └── main.go           # Entry point for the server
├── go.mod                # Go module file for managing dependencies
├── go.sum                # Go module checksum file
```

## Detailed Explanation of Files and Functions

### 1. `server/main.go`
This file is the entry point of the application, where the HTTP server is initialized and the main scraping logic resides. It exposes an HTTP endpoint for the API.

- **Functions:**
    - `main()`: Starts the HTTP server on port 8080 and listens for requests. The `handleRequest` function is registered as the handler for incoming requests.
    - `handleRequest(w http.ResponseWriter, r *http.Request)`: This function handles incoming requests. It reads the request body, decodes the JSON into a `RequestBody` struct, and calls `scrapeWithChromedp()` to scrape the specified URL. The scraped data is returned as a JSON response.
    - `scrapeWithChromedp(url string, body models.RequestBody) (models.ResponseBody, error)`: This function uses `chromedp` to scrape the web page content. It retrieves the page's title, HTML content, and then uses helper functions to extract links, paragraphs, images, headers, and emails based on the request options.
    - Helper functions within `main.go`:
        - `extractLinks(htmlContent string) []string`: Extracts all the links (`<a href="...">`) from the HTML content.
        - `extractParagraphs(htmlContent string) []string`: Extracts all the paragraphs (`<p>...</p>`) from the HTML content.
        - `extractHeaders(htmlContent string) []string`: Extracts all headers (`<h1>` to `<h6>`) from the HTML content.
        - `extractEmails(htmlContent string, emailRegex *regexp.Regexp) []string`: Uses a regex to extract email addresses from the HTML content.

### 2. `server/models/models.go`
This file defines the request and response data structures that the server uses to communicate with the client.

- **Structs:**
    - `RequestBody`: Represents the structure of the incoming request. It contains the following fields:
        - `Url`: The URL to scrape.
        - `Link`: A boolean to indicate if links only should be scraped.
        - `Paragraph`: A boolean to indicate if only paragraphs should be scraped.
        - `Image`: A boolean to indicate if only images should be scraped.
        - `Header`: A boolean to indicate if only headers should be scraped.
        - `Email`: A boolean to indicate if only emails should be scraped.
        - `All`: A boolean to indicate if all available information (links, paragraphs, headers, emails, images) should be scraped.
    - `ResponseBody`: Represents the structure of the server's response. It contains the following fields:
        - `Title`: The title of the scraped web page.
        - `Links`: A list of all links found on the page.
        - `Paragraphs`: A list of all paragraphs found on the page.
        - `Images`: A list of all image sources found on the page.
        - `Headers`: A list of all headers found on the page.
        - `Emails`: A list of all email addresses found on the page.

### 3. `server/helpers/extractImages.go`
This file contains helper functions used to process the scraped data.

- **Functions:**
    - `ExtractImages(htmlContent string, body models.RequestBody) []string`: This function takes the HTML content of the page and extracts all the image sources (`<img src="...">`). It returns a list of image URLs, and if the `src` attribute is a relative path, it prepends the base URL (from `RequestBody.Url`).

### 4. `server/helpers/extractHeaders.go`
This file contains helper functions used to process the scraped data.

- **Functions:**
    - `ExtractHeaders(htmlContent string) []string`: This function takes the HTML content of the page and extracts all the headers (`<h1>` to `<h6>`). It returns a list of header text.

### 5. `server/helpers/extractEmails.go`
This file contains helper functions used to process the scraped data.

- **Functions:**
    - `ExtractEmails(htmlContent string, emailRegex *regexp.Regexp) []string`: This function takes the HTML content of the page and extracts all the email addresses using a regex. It returns a list of email addresses.

### 6. `server/helpers/extractLinks.go`
This file contains helper functions used to process the scraped data.

- **Functions:**
    - `ExtractLinks(htmlContent string) []string`: This function takes the HTML content of the page and extracts all the links (`<a href="...">`). It returns a list of link URLs.

### 7. `server/helpers/extractParagraphs.go`
This file contains helper functions used to process the scraped data.

- **Functions:**
    - `ExtractParagraphs(htmlContent string) []string`: This function takes the HTML content of the page and extracts all the paragraphs (`<p>...</p>`). It returns a list of paragraph text.



### Dependencies
- **chromedp**: A headless browser package used to scrape dynamic content. Installed via:
    ```bash
    go get -u github.com/chromedp/chromedp
    ```
- **goquery**: Used for parsing and querying HTML documents. Installed via:
    ```bash
    go get -u github.com/PuerkitoBio/goquery
    ```

## How to Run

### Prerequisites
Ensure you have Go installed. You can check by running:
```bash
go version
```

### Steps to Run:
1. **Clone the repository**:
    ```bash
    git clone https://github.com/AnasAitZouinet/scrapper.git
    cd scrapper
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

3. **Run the application**:
    ```bash
    go run server/main.go
    ```

    This will start the server on `localhost:8080`.

### Example Request:
Once the server is running, you can use `curl` or Postman to make a request.

#### Curl Example:
```bash
curl -X POST "http://localhost:8080" -d '{"username":"test", "url":"https://example.com", "all":true}' -H "Content-Type: application/json"
```

This request will scrape the page at `https://example.com`, and return all the links, paragraphs, images, headers, and emails found on the page.

### Expected Response:
```json
{
    "title": "Example Domain",
    "links": ["https://www.iana.org/domains/example"],
    "paragraphs": ["This domain is for use in illustrative examples in documents."],
    "images": ["https://example.com/image.jpg"],
    "headers": ["Example Domain"],
    "emails": ["contact@example.com"]
}
```

### Notes:
- If `all: true` is specified in the request body, the server will extract all available content (links, paragraphs, images, headers, and emails). If `all` is false, only the content specified by `Link`, `Paragraph`, `Image`, `Header`, and `Email` flags will be extracted.
- This project is set up to scrape static and dynamic websites using a headless browser for full-page rendering.

### Troubleshooting:
- **Function not found**: Ensure that the function names are capitalized correctly when they are exported across packages.
- **Cannot access site content**: Some websites might block scraping requests from headless browsers. Ensure you are not blocked by firewalls or other restrictions.

---
