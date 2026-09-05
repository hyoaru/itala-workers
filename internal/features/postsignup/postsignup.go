package postsignup

import (
	workerport "github.com/hyoaru/itala-workers/internal/features/postsignup/application/port/worker"
	workeradapter "github.com/hyoaru/itala-workers/internal/features/postsignup/infrastructure/adapters/worker"
	"github.com/hyoaru/itala-workers/internal/shared/infrastructure/external/discordwebhookclient"
)

type (
	PostSignUpEvent  = workerport.PostSignUpEvent
	PostSignUpWorker = workerport.PostSignUpWorker
)

func NewDiscordPostSignUpWorker(webhookClient discordwebhookclient.DiscordWebhookClient) PostSignUpWorker {
	w := workeradapter.NewDiscordPostSignUpWorker(webhookClient)
	return workeradapter.NewDecoratedPostSignUpWorker(w)
}
