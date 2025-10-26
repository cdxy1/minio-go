package consumer

import (
	"log/slog"

	"github.com/cdxy1/minio-go/internal/config"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Handler interface {
	HandleMessage(msg []byte, offset kafka.Offset) error
}

type Consumer struct {
	consumer *kafka.Consumer
	handler  Handler
	stop     bool
	logger   *slog.Logger
}

func NewConsumer(handler Handler, logger *slog.Logger) (*Consumer, error) {
	cfg := config.GetConfig()

	kafkaCfg := kafka.ConfigMap{
		"bootstrap.servers":        cfg.Kafka.Host,
		"group.id":                 cfg.Kafka.Group,
		"session.timeout.ms":       cfg.Kafka.Timeout,
		"enable.auto.offset.store": cfg.Kafka.OffsetStore,
		"enable.auto.commit":       cfg.Kafka.AutoCommit,
		"auto.commit.interval.ms":  cfg.Kafka.CommitInterval,
	}

	c, err := kafka.NewConsumer(&kafkaCfg)
	if err != nil {
		logger.Error("Failed to create Kafka consumer", "error", err)
		return nil, err
	}

	if err := c.Subscribe(cfg.Kafka.Topic, nil); err != nil {
		logger.Error("Failed to subscribe to topic", "topic", cfg.Kafka.Topic, "error", err)
		return nil, err
	}

	logger.Info("Kafka consumer initialized",
		"host", cfg.Kafka.Host,
		"group", cfg.Kafka.Group,
		"topic", cfg.Kafka.Topic,
	)

	return &Consumer{consumer: c, handler: handler, logger: logger}, nil
}

func (c *Consumer) Start() {
	c.logger.Info("Kafka consumer started")

	for {
		if c.stop {
			c.logger.Info("Kafka consumer stopping")
			break
		}

		kafkaMsg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			c.logger.Error("Failed to read message", "error", err)
			continue
		}
		if kafkaMsg == nil {
			continue
		}

		c.logger.Info("Message received",
			"topic", *kafkaMsg.TopicPartition.Topic,
			"partition", kafkaMsg.TopicPartition.Partition,
			"offset", kafkaMsg.TopicPartition.Offset,
		)

		if err := c.handler.HandleMessage(kafkaMsg.Value, kafkaMsg.TopicPartition.Offset); err != nil {
			c.logger.Error("Handler failed",
				"offset", kafkaMsg.TopicPartition.Offset,
				"error", err,
			)
			continue
		}

		if _, err := c.consumer.StoreMessage(kafkaMsg); err != nil {
			c.logger.Error("Failed to store offset",
				"offset", kafkaMsg.TopicPartition.Offset,
				"error", err,
			)
			continue
		}

		c.logger.Info("Message processed successfully",
			"offset", kafkaMsg.TopicPartition.Offset,
		)
	}

	c.logger.Info("Kafka consumer stopped")
}

func (c *Consumer) Stop() error {
	c.logger.Info("Stopping Kafka consumer")

	c.stop = true
	if _, err := c.consumer.Commit(); err != nil {
		c.logger.Error("Failed to commit offsets on stop", "error", err)
		return err
	}

	if err := c.consumer.Close(); err != nil {
		c.logger.Error("Failed to close Kafka consumer", "error", err)
		return err
	}

	c.logger.Info("Kafka consumer closed successfully")
	return nil
}
