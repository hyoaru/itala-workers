package main

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hyoaru/itala-workers/internal/features/postsignup"
	"github.com/hyoaru/itala-workers/internal/shared/infrastructure/external/discordwebhookclient"
)

type Handler struct {
	worker      postsignup.PostSignUpWorker
	environment string
}

func (h *Handler) Handle(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	if err := h.worker.Execute(ctx, postsignup.PostSignUpEvent{
		Environment: h.environment,
		UserID:      event.Request.UserAttributes["sub"],
		Email:       event.Request.UserAttributes["email"],
		FirstName:   event.Request.UserAttributes["custom:first_name"],
		LastName:    event.Request.UserAttributes["custom:last_name"],
		OccurredAt:  time.Now(),
	}); err != nil {
		return event, err
	}

	return event, nil
}

func main() {
	environment := os.Getenv("ENVIRONMENT")
	discordWebhookURL := os.Getenv("DISCORD_WEBHOOK_URL")

	webhookClient := discordwebhookclient.NewHTTPDiscordWebhookClient(discordWebhookURL)
	worker := postsignup.NewDiscordPostSignUpWorker(webhookClient)
	handler := &Handler{worker, environment}

	lambda.Start(handler.Handle)
}
