package events

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Kafka interface {
	sendMessage(msg kafka.Message) error
	RunConsumers(ctx context.Context, topics []TopicConfig) error
}

type broker struct {
	Writer *kafka.Writer
	cfg    *config.Config
	log    *logrus.Logger
}

type InternalEvent struct {
	EventType string          `json:"event_type"`
	Data      json.RawMessage `json:"data"`
}

// TopicConfig теперь требует другой сигнатуры Callback:
type TopicConfig struct {
	Topic    string
	Callback func(ctx context.Context, m kafka.Message, evt InternalEvent) error
}

func NewBroker(cfg *config.Config, log *logrus.Logger) Kafka {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(cfg.Kafka.Brokers...),
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    1,
		BatchTimeout: 0,
		Async:        true,
	}

	return &broker{
		Writer: writer,
		cfg:    cfg,
		log:    log,
	}
}

func (b *broker) sendMessage(msg kafka.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := b.Writer.WriteMessages(ctx, msg); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

func (b *broker) RunConsumers(ctx context.Context, topics []TopicConfig) error {
	for _, t := range topics {
		tc := t

		go func() {
			r := kafka.NewReader(kafka.ReaderConfig{
				Brokers:  b.cfg.Kafka.Brokers,
				Topic:    tc.Topic,
				MinBytes: 1,
				MaxBytes: 10e6,
			})
			defer r.Close()

			for {
				m, err := r.ReadMessage(ctx)
				if err != nil {
					if ctx.Err() != nil {
						return
					}
					b.log.WithField("kafka", err).Errorf("Error reading message from topic %s", tc.Topic)
					return
				}

				var ie InternalEvent
				if err := json.Unmarshal(m.Value, &ie); err != nil {
					b.log.WithField("kafka", err).Errorf("Error unmarshalling InternalEvent")
					continue
				}

				if cbErr := tc.Callback(ctx, m, ie); cbErr != nil {
					b.log.WithField("kafka", cbErr).Errorf("Error processing message from topic %s", tc.Topic)
				}
			}
		}()
	}

	return nil
}
