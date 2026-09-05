package discordwebhookclient

type EmbedField struct {
	Name   string
	Value  string
	Inline bool
}

type EmbedFooter struct {
	Text string
}

type Embed struct {
	Title       string
	Description string
	Color       int
	Fields      []EmbedField
	Footer      EmbedFooter
}

type DiscordWebhookClient interface {
	Send(content string, embeds []Embed) error
}
