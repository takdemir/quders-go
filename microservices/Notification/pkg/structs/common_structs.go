package structs

import "encoding/json"

type EmailFeatures struct {
	From                        map[string]string   `json:"from"`
	Tos                         []map[string]string `json:"tos"`
	Subject                     string              `json:"subject"`
	PlainText                   string              `json:"plainText"`
	HtmlContent                 string              `json:"htmlContent"`
	HtmlContentTemplateFileName string              `json:"htmlContentTemplateFileName"`
	HtmlTemplateData            map[string]string   `json:"htmlTemplateData"`
}

type SlackFeatures struct {
	Color      string      `json:"color"`
	AuthorName string      `json:"authorName"`
	AuthorLink string      `json:"authorLink"`
	AuthorIcon string      `json:"authorIcon"`
	Text       string      `json:"text"`
	Ts         json.Number `json:"ts"`
}

type TeamsFeatures struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	ThemeColor string `json:"themeColor"`
}
