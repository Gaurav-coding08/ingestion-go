package kafka_test

import (
	"testing"

	internalKafka "github.com/Gaurav-coding08/ingestion-go/cmd/kafka"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error {
	args := m.Called(msg, deliveryChan)

	go func() {
		deliveryChan <- &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     msg.TopicPartition.Topic,
				Partition: 0,
				Offset:    1,
			},
			Value: msg.Value,
		}
		close(deliveryChan)
	}()

	return args.Error(0)
}

func TestProduceMessage_Success(t *testing.T) {
	mockProducer := new(MockProducer)

	topic := "test_topic"
	eventType := "stock.update"
	payload := []byte(`{"id": 1, "name":"tes_stock", "price": 200.50}`)

	mockProducer.On("Produce", mock.Anything, mock.Anything).Return(nil)

	err := internalKafka.ProduceMessage(mockProducer, topic, eventType, payload)

	assert.NoError(t, err, "Expected no error in message production")
	mockProducer.AssertCalled(t, "Produce", mock.Anything, mock.Anything)
}

// Test Passes When Kafka Producer Returns an Error
func TestProduceMessage_Failure(t *testing.T) {
	mockProducer := new(MockProducer)

	topic := "test_topic"
	eventType := "stock.update"
	payload := []byte(`{"id": 1, "price": 200.50}`)

	mockProducer.On("Produce", mock.Anything, mock.Anything).Return(assert.AnError)

	err := internalKafka.ProduceMessage(mockProducer, topic, eventType, payload)

	assert.Error(t, err, "Expected an error in message production")
	mockProducer.AssertCalled(t, "Produce", mock.Anything, mock.Anything)
}
