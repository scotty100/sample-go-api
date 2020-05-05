package service

import "context"
import "github.com/BenefexLtd/onehub-go-base/pkg/messaging"

type MessagePublisher interface {
	Publish(ctx context.Context, payload messaging.OneHubEvent, messageType string) (string, error)
}