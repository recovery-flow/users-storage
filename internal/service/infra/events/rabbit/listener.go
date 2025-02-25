package rabbit

import (
	"context"

	"github.com/recovery-flow/rerabbit"
	"github.com/recovery-flow/users-storage/internal/service"
	"github.com/recovery-flow/users-storage/internal/service/domain/callbacks"
	"github.com/recovery-flow/users-storage/internal/service/infra/events/rabbit/amqpconfig"
	"github.com/streadway/amqp"
)

var eventDispatcher = map[string]func(context.Context, *service.Service, []byte) error{
	amqpconfig.AccountUpdateRoleKey: callbacks.AccountUpdateRole,
	amqpconfig.AccountCreateKey:     callbacks.AccountCreate,
}

func Listener(ctx context.Context, svc *service.Service) {
	rabbitWorker, err := rerabbit.NewBroker(svc.Config.Rabbit.URL)
	if err != nil {
		svc.Log.Errorf("Failed to connect to RabbitMQ: %v", err)
		<-ctx.Done()
		return
	}

	go func() {
		<-ctx.Done()
		svc.Log.Info("Shutting down RabbitMQ connection...")
		rabbitWorker.GracefulShutdown(svc.Log)
	}()

	opts := rerabbit.ConsumeOptions{
		QueueName:   amqpconfig.AccountUsersStorageQ, // Одна очередь
		ConsumerTag: "",
		AutoAck:     false,
		Exclusive:   false,
		NoLocal:     false,
		NoWait:      false,
		Args:        nil,
	}

	//type QueueConfig struct {
	//	QueueName  string
	//	RoutingKey string
	//	Callback   func(context.Context, *service.Service, []byte) error
	//}

	//queues := []QueueConfig{
	//	{
	//		QueueName: amqpconfig.AccountUsersStorageQ,
	//		Callback:  callbacks.AccountUpdateRole,
	//	},
	//	{
	//		QueueName: amqpconfig.AccountUsersStorageQ,
	//		Callback:  callbacks.AccountCreate,
	//	},
	//}

	//for _, qc := range queues {
	//	qc := qc // захватываем локальную копию, чтобы избежать гонок
	//	go func(qc QueueConfig) {
	//		opts := rerabbit.ConsumeOptions{
	//			QueueName:   qc.QueueName,
	//			ConsumerTag: "",
	//			AutoAck:     false,
	//			Exclusive:   false,
	//			NoLocal:     false,
	//			NoWait:      false,
	//			Args:        nil,
	//		}
	//
	//		err := rabbitWorker.Consume(ctx, opts, func(ctx context.Context, d amqp.Delivery) {
	//			if err := qc.Callback(ctx, svc, d.Body); err != nil {
	//				svc.Log.Errorf("Error processing message from queue %s: %v", qc.QueueName, err)
	//				if nackErr := d.Nack(false, true); nackErr != nil {
	//					svc.Log.Errorf("Failed to Nack message: %v", nackErr)
	//				}
	//			} else {
	//				if ackErr := d.Ack(false); ackErr != nil {
	//					svc.Log.Errorf("Failed to Ack message: %v", ackErr)
	//				}
	//			}
	//		})
	//		if err != nil {
	//			svc.Log.Errorf("Error consuming from queue %s: %v", qc.QueueName, err)
	//		}
	//	}(qc)
	//}

	err = rabbitWorker.Consume(ctx, opts, func(ctx context.Context, d amqp.Delivery) {
		handler, exists := eventDispatcher[d.RoutingKey] // Ищем обработчик по ключу
		if !exists {
			svc.Log.Warnf("Received unknown event: %s", d.RoutingKey)
			_ = d.Nack(false, false) // Не отправляем в повторную обработку
			return
		}

		if err := handler(ctx, svc, d.Body); err != nil {
			svc.Log.Errorf("Error processing event %s: %v", d.RoutingKey, err)
			if nackErr := d.Nack(false, true); nackErr != nil {
				svc.Log.Errorf("Failed to Nack message: %v", nackErr)
			}
		} else {
			if ackErr := d.Ack(false); ackErr != nil {
				svc.Log.Errorf("Failed to Ack message: %v", ackErr)
			}
		}
	})

	if err != nil {
		svc.Log.Errorf("Error consuming from queue %s: %v", amqpconfig.AccountUsersStorageQ, err)
	}

	<-ctx.Done()
	svc.Log.Info("Listener events [RabbitMQ] stopped")
}
