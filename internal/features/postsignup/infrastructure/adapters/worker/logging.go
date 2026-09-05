package postsignup

import (
	"context"

	port "github.com/hyoaru/itala-workers/internal/features/postsignup/application/port/worker"
	"github.com/hyoaru/itala-workers/internal/shared/infrastructure/logger"
)

type LoggingPostSignUpWorker struct {
	inner port.PostSignUpWorker
}

func NewLoggingPostSignUpWorker(inner port.PostSignUpWorker) *LoggingPostSignUpWorker {
	return &LoggingPostSignUpWorker{inner: inner}
}

func (w *LoggingPostSignUpWorker) Execute(ctx context.Context, event port.PostSignUpEvent) error {
	if err := w.inner.Execute(ctx, event); err != nil {
		logger.Error("Failed to execute post sign-up event", "error", err)
		return err
	}

	logger.Info("Executed post sign-up event", "user", event.UserID)
	return nil
}
