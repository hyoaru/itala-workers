package discordwebhookclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPDiscordWebhookClient struct {
	url        string
	httpClient *http.Client
}

func NewHTTPDiscordWebhookClient(url string) DiscordWebhookClient {
	return &HTTPDiscordWebhookClient{url: url, httpClient: http.DefaultClient}
}

type embedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type embedFooter struct {
	Text string `json:"text"`
}

type embed struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Color       int          `json:"color"`
	Fields      []embedField `json:"fields"`
	Footer      embedFooter  `json:"footer"`
}

func toEmbedDTO(e Embed) embed {
	fields := make([]embedField, 0, len(e.Fields))

	for _, field := range e.Fields {
		fields = append(fields, embedField(field))
	}

	return embed{
		Title:       e.Title,
		Description: e.Description,
		Color:       e.Color,
		Fields:      fields,
		Footer:      embedFooter(e.Footer),
	}
}

type sendRequest struct {
	Content string  `json:"content,omitempty"`
	Embeds  []embed `json:"embeds"`
}

func (c *HTTPDiscordWebhookClient) Send(content string, embeds []Embed) error {
	embedDTOs := make([]embed, 0, len(embeds))

	for _, e := range embeds {
		embedDTOs = append(embedDTOs, toEmbedDTO(e))
	}

	request := sendRequest{Content: content, Embeds: embedDTOs}

	body, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("marshal discord webhook payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create discord webhook request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("send discord webhook request: %w", err)
	}

	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("discord webhook returned status %s", resp.Status)
	}

	return nil
}
