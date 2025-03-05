package listener

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/recovery-flow/users-storage/internal/service"
	"github.com/recovery-flow/users-storage/internal/service/domain/callbacks"
	"github.com/recovery-flow/users-storage/internal/service/infra/events"
	"github.com/recovery-flow/users-storage/internal/service/infra/events/evebody"
	"github.com/recovery-flow/users-storage/internal/service/infra/events/kafig"
	"github.com/segmentio/kafka-go"
)

func Listener(ctx context.Context, svc *service.Service) {
	broker := events.NewBroker(svc.Config, svc.Log)

	topics := []events.TopicConfig{
		{
			Topic: kafig.AccountsTopic,
			Callback: func(ctx context.Context, m kafka.Message, ie events.InternalEvent) error {
				// Смотрим event_type
				switch ie.EventType {
				case "account_created":
					// Парсим поле Data как AccountCreated
					var ac evebody.AccountCreated
					if err := json.Unmarshal(ie.Data, &ac); err != nil {
						return fmt.Errorf("failed to unmarshal AccountCreated: %w", err)
					}
					return callbacks.AccountCreate(ctx, svc, ac)

				case "account_role_updated":
					return nil

				default:
					svc.Log.WithField("kafka", m).Warn("Unknown event type")
				}
				return nil
			},
		},
	}

	if err := broker.RunConsumers(ctx, topics); err != nil {
		svc.Log.Errorf("Error running consumers: %v", err)
	}

	<-ctx.Done()
	svc.Log.Info("Kafka listener stopped")
}
