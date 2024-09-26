package models


type RequestBody struct {
	Url       string `json:"url"`
	Link      bool   `json:"link"`
	Paragraph bool   `json:"paragraph"`
	Image     bool   `json:"image"`
	Header    bool   `json:"header"`
	Email     bool   `json:"email"`
	All       bool   `json:"all"`
}

type ResponseBody struct {
	Title      string   `json:"title"`
	Links      []string `json:"links"`
	Paragraphs []string `json:"paragraphs"`
	Images     []string `json:"images"`
	Headers    []string `json:"headers"`
	Emails     []string `json:"emails"`
}