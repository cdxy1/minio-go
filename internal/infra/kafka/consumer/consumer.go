package consumer

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type Handler interface {
	HandleMessage(msg []byte, offset kafka.Offset) error
}

type Consumer struct {
	consumer *kafka.Consumer
	handler  Handler
	stop     bool
}

func NewConsumer(handler Handler) (*Consumer, error) {
	kafkaCfg := kafka.ConfigMap{
		"bootstrap.servers":        "",
		"group.id":                 "metadata-group",
		"session.timeout.ms":       5000,
		"enable.auto.offset.store": false,
		"enable.auto.commit":       true,
		"auto.commit.interval.ms":  5000,
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
