package services

import (
	"encoding/json"
	"log"

	internalKafka "github.com/Gaurav-coding08/ingestion-go/cmd/kafka"
	v1 "github.com/Gaurav-coding08/ingestion-go/pkg/client"
	common "github.com/Gaurav-coding08/ingestion-go/pkg/common"
		"github.com/Gaurav-coding08/ingestion-go/cmd/kafka"
)

type service struct {
	producer kafka.KafkaProducer
}

func New(producer kafka.KafkaProducer) *service {
	return &service{
		producer: producer,
	}
}

func (s *service) UpdateStockPrice(
	updateStockReq v1.UpdateStockPrice,
) error {

	eventPayload := common.StockUpdate{
		ID:    updateStockReq.ID,
		Price: updateStockReq.Price,
		Name:  updateStockReq.Name,
	}

	eventBytes, err := json.Marshal(eventPayload)
	if err != nil {
		log.Println("Failed to marshal stock update payload:", err)
		return err
	}

	err = internalKafka.ProduceMessage(s.producer, common.LiveUpdatesTopic, string(common.EventStockUpdate), eventBytes)
	if err != nil {
		log.Println("Error sending stock update to Kafka:", err)

		return err
	}

	return nil
}
