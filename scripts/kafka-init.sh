#!/bin/sh

echo "Waiting for Kafka..."
sleep 2

/opt/kafka/bin/kafka-topics.sh --create \
  --bootstrap-server localhost:9092 \
  --replication-factor 1 \
  --partitions 3 \
  --topic my-topic
