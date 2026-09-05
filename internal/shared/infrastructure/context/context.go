package context

import (
	"context"
	"errors"
)

type contextKey string

const environmentContextKey contextKey = "environment"

func WithEnvironment(ctx context.Context, environment string) context.Context {
	return context.WithValue(ctx, environmentContextKey, environment)
}

func EnvironmentFromContext(ctx context.Context) (string, error) {
	environment, ok := ctx.Value(environmentContextKey).(string)
	if !ok {
		return "", errors.New("environment not found in context")
	}

	return environment, nil
}
