package common

import (
	"fmt"
)

var EventModelMap = map[EventType]interface{}{
	EventStockUpdate: StockUpdate{},
}

func GetEventModel(eventType EventType) (interface{}, error) {
	model, exists := EventModelMap[eventType]
	if !exists {
		return nil, fmt.Errorf("❌ Unknown event type: %s", eventType)
	}
	return model, nil
}
