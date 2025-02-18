package common

type EventType string

// events can be implmented according to use-case
const (
	EventStockUpdate    EventType = "stock.update"
	EventOrderCreated   EventType = "order.created"
	EventUserRegistered EventType = "user.registered"
)
