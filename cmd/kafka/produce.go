package kafka

import (
	"encoding/json"
	"log"

	common "github.com/Gaurav-coding08/ingestion-go/pkg/common"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

type KafkaProducer interface {
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
}

func ProduceMessage(producer KafkaProducer, topic string, eventType string, payload []byte) error {
	deliveryChan := make(chan kafka.Event, 1)

	kafkaMsg := common.KafkaMessage{
		ID:        uuid.New().String(),
		EventType: eventType,
		Payload:   payload,
	}

	// Convert KafkaMessage to JSON bytes
	kafkaMsgBytes, err := json.Marshal(kafkaMsg)
	if err != nil {
		log.Printf("❌ Failed to marshal KafkaMessage: %v", err)
		return err
	}

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          kafkaMsgBytes,
	}, deliveryChan)

	if err != nil {
		log.Printf("Failed to produce message: %v", err)
		return err
	}

	go func() {
		e := <-deliveryChan
		m, ok := e.(*kafka.Message)
		if !ok {
			log.Printf("❌ Unexpected event type")
			return
		}
		if m.TopicPartition.Error != nil {
			log.Printf("Message Delivery failed: %v", m.TopicPartition.Error)
		} else {
			log.Printf("Message delivered to topic %s [%d] at offset %d\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
		//close(deliveryChan)
	}()

	return nil
}
