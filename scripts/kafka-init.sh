#!/bin/sh

echo "Waiting for Kafka 10 seconds"
sleep 10

/opt/kafka/bin/kafka-topics.sh --create \
  --bootstrap-server localhost:9092 \
  --replication-factor 1 \
  --partitions 3 \
  --topic metadata-topic
