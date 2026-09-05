package presignup

import (
	"context"
	"time"
)

type PreSignUpEvent struct {
	Environment string
	Email       string
	FirstName   string
	LastName    string
	OccurredAt  time.Time
}

type PreSignUpWorker interface {
	Execute(ctx context.Context, event PreSignUpEvent) error
}
