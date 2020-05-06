package kafka

import (
	"time"

	"github.com/Shopify/sarama"
)

// KafkaHelper ...
type KafkaHelper struct {
	Host     string
	Username string
	Password string
}

// Main ...
func (t *KafkaHelper) config() *sarama.Config {
	config := sarama.NewConfig()
	// config.ClientID = "go-kafka-consumer"
	config.Consumer.Return.Errors = true
	config.Producer.Return.Successes = true
	config.Net.WriteTimeout = 5 * time.Second
	config.Producer.Retry.Max = 5

	if t.Username != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = t.Username
		config.Net.SASL.Password = t.Password
	}
	return config
}
