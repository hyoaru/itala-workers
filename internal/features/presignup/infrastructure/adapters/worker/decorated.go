package presignup

import (
	"context"

	port "github.com/hyoaru/itala-workers/internal/features/presignup/application/port/worker"
)

type DecoratedPreSignUpWorker struct {
	inner port.PreSignUpWorker
}

func NewDecoratedPreSignUpWorker(inner port.PreSignUpWorker) *DecoratedPreSignUpWorker {
	return &DecoratedPreSignUpWorker{inner: NewLoggingPreSignUpWorker(inner)}
}

func (w *DecoratedPreSignUpWorker) Execute(ctx context.Context, event port.PreSignUpEvent) error {
	return w.inner.Execute(ctx, event)
}
