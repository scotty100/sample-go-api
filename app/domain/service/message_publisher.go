package service

import "context"
import "github.com/BenefexLtd/departments-api-refactor/app/utl/messaging"

type MessagePublisher interface {
	Publish(ctx context.Context, payload messaging.OneHubEvent, messageType string) (string, error)
}