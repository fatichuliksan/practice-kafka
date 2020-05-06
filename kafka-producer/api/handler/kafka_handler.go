package handler

import (
	"kafka-example/helper"
	"time"

	"github.com/Shopify/sarama"
	"github.com/labstack/echo"
)

// NewKafkaHandler ...
type NewKafkaHandler struct {
	Helper helper.NewHelper
}

// Test ...
func (t *NewKafkaHandler) Test(c echo.Context) error {
	msg := c.QueryParam("msg")
	topic := c.QueryParam("topic")
	producer := t.Helper.Kafka.Producer()

	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder("test"),
		Value: sarama.StringEncoder(msg),
		// Headers: []RecordHeader,
		// Metadata: interface{},
		// Offset int64
		// Partition int32
		Timestamp: time.Now(),
	}

	producer.SendMessage(kafkaMsg)

	return t.Helper.Response.SendSuccess(c, "Test Kafka Handler", nil)
}
