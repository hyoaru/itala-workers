package presignup

import (
	workerport "github.com/hyoaru/itala-workers/internal/features/presignup/application/port/worker"
	workeradapter "github.com/hyoaru/itala-workers/internal/features/presignup/infrastructure/adapters/worker"
	"github.com/hyoaru/itala-workers/internal/shared/infrastructure/external/discordwebhookclient"
)

type (
	PreSignUpEvent  = workerport.PreSignUpEvent
	PreSignUpWorker = workerport.PreSignUpWorker
)

func NewDiscordPreSignUpWorker(webhookClient discordwebhookclient.DiscordWebhookClient) PreSignUpWorker {
	w := workeradapter.NewDiscordPreSignUpWorker(webhookClient)
	return workeradapter.NewDecoratedPreSignUpWorker(w)
}
