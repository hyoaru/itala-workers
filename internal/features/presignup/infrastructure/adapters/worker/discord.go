package presignup

import (
	"context"
	"fmt"
	"strings"

	port "github.com/hyoaru/itala-workers/internal/features/presignup/application/port/worker"
	"github.com/hyoaru/itala-workers/internal/shared/infrastructure/external/discordwebhookclient"
)

type DiscordPreSignUpWorker struct {
	webhookClient discordwebhookclient.DiscordWebhookClient
}

func NewDiscordPreSignUpWorker(webhookClient discordwebhookclient.DiscordWebhookClient) *DiscordPreSignUpWorker {
	return &DiscordPreSignUpWorker{webhookClient: webhookClient}
}

func (w *DiscordPreSignUpWorker) Execute(ctx context.Context, event port.PreSignUpEvent) error {
	embed := discordwebhookclient.Embed{
		Title:       fmt.Sprintf("%s: Pre-confirmation Sign Up", strings.ToUpper(event.Environment)),
		Description: "A user is attempting to sign up for itala!",
		Color:       16776960,
		Fields: []discordwebhookclient.EmbedField{
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
