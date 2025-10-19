package producer

import (
	"errors"

	"github.com/cdxy1/go-file-storage/internal/config"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var errUnknownType = errors.New("unknown type")

type Producer struct {
	producer *kafka.Producer
}

func NewProducer() (*Producer, error) {
	cfg := config.GetConfig()

	conf := kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Host,
	}
	p, err := kafka.NewProducer(&conf)
	if err != nil {
		return nil, err
	}

	return &Producer{producer: p}, nil
}

func (p *Producer) Produce(msg, topic string) error {
	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(msg),
		Key:   nil,
	}
	kafkaChan := make(chan kafka.Event)
	if err := p.producer.Produce(kafkaMsg, kafkaChan); err != nil {
		return err
	}
	e := <-kafkaChan

	switch ev := e.(type) {
	case *kafka.Message:
		return nil
	case *kafka.Error:
		return ev
	default:
		return errUnknownType
	}
}

func (p *Producer) Close() {
	p.producer.Flush(5000)
	p.producer.Close()
}
