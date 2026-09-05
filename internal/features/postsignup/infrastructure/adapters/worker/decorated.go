package postsignup

import (
	"context"

	port "github.com/hyoaru/itala-workers/internal/features/postsignup/application/port/worker"
)

type DecoratedPostSignUpWorker struct {
	inner port.PostSignUpWorker
}

func NewDecoratedPostSignUpWorker(inner port.PostSignUpWorker) *DecoratedPostSignUpWorker {
	return &DecoratedPostSignUpWorker{inner: NewLoggingPostSignUpWorker(inner)}
}

func (w *DecoratedPostSignUpWorker) Execute(ctx context.Context, event port.PostSignUpEvent) error {
	return w.inner.Execute(ctx, event)
}
