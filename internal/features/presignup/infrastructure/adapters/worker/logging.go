package presignup

import (
	"context"

	port "github.com/hyoaru/itala-workers/internal/features/presignup/application/port/worker"
	"github.com/hyoaru/itala-workers/internal/shared/infrastructure/logger"
)

type LoggingPreSignUpWorker struct {
	inner port.PreSignUpWorker
}

func NewLoggingPreSignUpWorker(inner port.PreSignUpWorker) *LoggingPreSignUpWorker {
	return &LoggingPreSignUpWorker{inner: inner}
}

func (w *LoggingPreSignUpWorker) Execute(ctx context.Context, event port.PreSignUpEvent) error {
	logger.Debug("Executing pre sign-up event", "email", event.Email)
	if err := w.inner.Execute(ctx, event); err != nil {
		logger.Error("Failed to execute pre sign-up event", "error", err)
		return err
	}

	logger.Debug("Executed pre sign-up event", "email", event.Email)
	return nil
}
