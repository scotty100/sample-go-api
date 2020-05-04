package messaging

import (
	"context"
	"github.com/BenefexLtd/departments-api-refactor/app/utl/messaging"
	uuid "github.com/satori/go.uuid"
)

// this would come from a gcp pub sub messaging package that wraps around the provided publisher methods and adds our headers etc

// PublisherConfig to be provided by the consumer.
type Publisher struct {
	Topic string
}

// Publish message to pubsub
func (publisher *Publisher) Publish(ctx context.Context, payload messaging.OneHubEvent, messageType string) (string, error) {

	attributes := make(map[string]string, 1)
	attributes["message_type"] = messageType
	//data, err := json.Marshal(payload)
	//if err != nil {
	//	return ``, err
	//}
	//message := &pubsub.Message{
	//	Data:       data,
	//	Attributes: attributes,
	//}
	//response := publisher.topic.Publish(ctx, message)
	//return response.Get(ctx)

	return uuid.NewV4().String(), nil
}
