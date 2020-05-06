package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// Producer ..
func (t *KafkaHelper) Producer() sarama.SyncProducer {
	kafkaConfig := t.config()
	producers, err := sarama.NewSyncProducer([]string{t.Host}, kafkaConfig)
	if err != nil {
		fmt.Println(fmt.Sprintf("Unable to create kafka producer got error %v", err))
		return nil
	}
	// defer func() {
	// 	if err := producers.Close(); err != nil {
	// 		fmt.Println(fmt.Sprintf("Unable to stop kafka producer: %v", err))
	// 	}
	// }()
	return producers
}
