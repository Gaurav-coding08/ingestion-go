package common

import (
	"encoding/json"
)

type KafkaMessage struct {
	ID        string          `json:"id"`
	EventType string          `json:"event_type"` // type of event
	Payload   json.RawMessage `json:"payload"`    // Stores raw JSON data, so that it is general and can be unmarshalled in any consumer according to topic
}

// without omitempty it will throw error if field not provided in the kafka message,
type StockUpdate struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
