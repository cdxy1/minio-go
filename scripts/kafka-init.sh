#!/bin/sh

while ! nc -z kafka 9092; do
  echo "Waiting for Kafka..."
  sleep 2
done

bin/kafka-topics.sh --create \
  --bootstrap-server localhost:9092 \
  --replication-factor 1 \
  --partitions 3 \
  --topic my-topic
