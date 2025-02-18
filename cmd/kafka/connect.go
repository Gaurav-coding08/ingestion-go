package kafka

import (
	"log"

	"github.com/Gaurav-coding08/ingestion-go/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ConnectToProducer(cfg *config.AppConfig) (KafkaProducer, error) {
	broker := cfg.KafkaConfig.Broker

	// initialize based on any env and config accordingly.
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"acks":              "all",
	})
	if err != nil {
		log.Printf("Failed to create Kafka producer: %v", err)

		return nil, err
	}

	return producer, nil
}
