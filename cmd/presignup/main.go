package main

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hyoaru/itala-workers/internal/features/presignup"
	"github.com/hyoaru/itala-workers/internal/shared/infrastructure/external/discordwebhookclient"
)

type Handler struct {
	worker      presignup.PreSignUpWorker
	environment string
}

func (h *Handler) Handle(ctx context.Context, event events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignup, error) {
	if err := h.worker.Execute(ctx, presignup.PreSignUpEvent{
		Environment: h.environment,
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
	worker := presignup.NewDiscordPreSignUpWorker(webhookClient)
	handler := &Handler{worker, environment}

	lambda.Start(handler.Handle)
}
