package postsignup

import (
	"context"
	"time"
)

type PostSignUpEvent struct {
	Environment string
	UserID      string
	Email       string
	FirstName   string
	LastName    string
	OccurredAt  time.Time
}

type PostSignUpWorker interface {
	Execute(ctx context.Context, event PostSignUpEvent) error
}
