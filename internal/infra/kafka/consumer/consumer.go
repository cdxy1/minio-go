package consumer

import (
	"github.com/cdxy1/go-file-storage/internal/config"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Handler interface {
	HandleMessage(msg []byte, offset kafka.Offset) error
}

type Consumer struct {
	consumer *kafka.Consumer
	handler  Handler
	stop     bool
}

func NewConsumer(handler Handler) (*Consumer, error) {
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
		return nil, err
	}

	if err := c.Subscribe("some-topic", nil); err != nil {
		return nil, err
	}
	return &Consumer{consumer: c, handler: handler}, nil
}

func (c *Consumer) Start() {
	for {
		if c.stop {
			break
		}
		kafkaMsg, err := c.consumer.ReadMessage(7000)
		if err != nil {
			println(err.Error())
		}
		if kafkaMsg == nil {
			continue
		}

		if err := c.handler.HandleMessage(kafkaMsg.Value, kafkaMsg.TopicPartition.Offset); err != nil {
			println(err.Error())
			continue
		}

		if _, err := c.consumer.StoreMessage(kafkaMsg); err != nil {
			println(err.Error())
			continue
		}
	}
}

func (c *Consumer) Stop() error {
	c.stop = true
	if _, err := c.consumer.Commit(); err != nil {
		return err
	}
	return c.consumer.Close()
}
