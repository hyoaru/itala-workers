package postsignup

import (
	"context"
	"fmt"

	port "github.com/hyoaru/itala-workers/internal/features/postsignup/application/port/worker"
	"github.com/hyoaru/itala-workers/internal/shared/infrastructure/external/discordwebhookclient"
)

type DiscordPostSignUpWorker struct {
	webhookClient discordwebhookclient.DiscordWebhookClient
}

func NewDiscordPostSignUpWorker(webhookClient discordwebhookclient.DiscordWebhookClient) *DiscordPostSignUpWorker {
	return &DiscordPostSignUpWorker{webhookClient: webhookClient}
}

func (w *DiscordPostSignUpWorker) Execute(ctx context.Context, event port.PostSignUpEvent) error {
	embed := discordwebhookclient.Embed{
		Title:       fmt.Sprintf("%s: Post-confirmation Sign Up", event.Environment),
		Description: "A new user has joined itala!",
		Color:       14807354,
		Fields: []discordwebhookclient.EmbedField{
			{
				Name:   "User ID",
				Value:  event.UserID,
				Inline: false,
			},
			{
				Name:   "Email",
				Value:  event.Email,
				Inline: true,
			},
			{
				Name:   "Name",
				Value:  fmt.Sprintf("%s %s", event.FirstName, event.LastName),
				Inline: true,
			},
		},
		Footer: discordwebhookclient.EmbedFooter{
			Text: event.OccurredAt.Format("January 2, 2006"),
		},
	}

	if err := w.webhookClient.Send("", []discordwebhookclient.Embed{embed}); err != nil {
		return fmt.Errorf("send discord webhook: %w", err)
	}

	return nil
}
