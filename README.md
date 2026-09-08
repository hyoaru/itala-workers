# Itala Workers

AWS Lambda functions for the [Itala](https://github.com/hyoaru/itala-pwa) personal finance platform. Handles background tasks and event-driven workflows across the ecosystem. Part of the Itala ecosystem alongside the [API backend](https://github.com/hyoaru/itala-api), [PWA frontend](https://github.com/hyoaru/itala-pwa), and [AWS CDK infrastructure](https://github.com/hyoaru/itala-infrastructure).

## Architecture

Each worker is an isolated Lambda function under `cmd/`, following clean architecture with a decorator pattern for cross-cutting concerns like logging and notifications. Workers are event-driven and designed to be added incrementally as the platform grows.

### Platform Architecture

![Itala Infrastructure](docs/assets/Itala%20Infrastructure.png)

## Project Structure

```
itala-workers/
├── docs/assets/                # Architecture diagrams
├── cmd/
│   ├── presignup/main.go        # Lambda entrypoint (PreSignUp trigger)
│   └── postsignup/main.go       # Lambda entrypoint (PostConfirmation trigger)
├── internal/
│   ├── features/
│   │   ├── presignup/           # Pre-signup feature
│   │   │   ├── application/port/worker/  # PreSignUpWorker interface
│   │   │   └── infrastructure/adapters/worker/
│   │   │       ├── discord.go    # Discord notification logic
│   │   │       ├── decorated.go  # Decorator wrapper
│   │   │       └── logging.go    # Logging decorator
│   │   └── postsignup/          # Post-confirmation feature
│   │       ├── application/port/worker/  # PostSignUpWorker interface
│   │       └── infrastructure/adapters/worker/
│   │           ├── discord.go
│   │           ├── decorated.go
│   │           └── logging.go
│   └── shared/
│       └── infrastructure/
│           ├── external/discordwebhookclient/  # Discord webhook HTTP client
│           └── logger/                          # slog-based logger
└── go.mod
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `ENVIRONMENT` | Environment name (e.g. `staging`, `production`) — displayed in notification titles |
| `DISCORD_WEBHOOK_URL` | Discord webhook URL for sending notifications |

## Tech Stack

- **Go 1.26** — core language
- **aws-lambda-go** — Lambda runtime
- **Discord Webhook** — HTTP client for notifications

## Prerequisites

- Go 1.26+
- AWS credentials configured
- `ENVIRONMENT` and `DISCORD_WEBHOOK_URL` environment variables

## Workers

### Pre-SignUp Worker

**Trigger:** Cognito PreSignUp event (fires before user confirmation)

**Purpose:** Notifies Discord that a user is attempting to register.

**Does not block sign-up** — always returns the original event unchanged.

### Post-Confirmation Worker

**Trigger:** Cognito PostConfirmation event (fires after email verification)

**Purpose:** Notifies Discord that a new user has fully joined.

**Does not block sign-up** — always returns the original event unchanged.

## Deployment

### CI/CD Pipeline

Two GitHub Actions workflows, one per Lambda:

| Workflow | Branch | Environment |
|----------|--------|-------------|
| `cd-pre-confirmation-sign-up.yml` | `develop` → staging | staging |
| `cd-pre-confirmation-sign-up.yml` | `master` → production | production |
| `cd-post-confirmation-sign-up.yml` | `develop` → staging | staging |
| `cd-post-confirmation-sign-up.yml` | `master` → production | production |

Workflows are path-filtered to only run when relevant source files change.

### Pipeline Steps

1. **Build** — cross-compile Go binary for linux/arm64, package as `function.zip`
2. **Deploy** — assumes AWS role via OIDC, uploads to S3, updates Lambda function code
3. **Notify** — sends build status to Discord
