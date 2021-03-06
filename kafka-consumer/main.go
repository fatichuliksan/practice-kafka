package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

// KafkaHelper ...
type KafkaHelper struct {
	Host     string
	Username string
	Password string
}

func main() {
	kafkaNew := KafkaHelper{
		Host: "localhost:9092",
	}
	kafkaNew.Consumer()
}

// Main ...
func (t *KafkaHelper) config() *sarama.Config {
	config := sarama.NewConfig()
	config.ClientID = "kafka-consumer"
	config.Consumer.Return.Errors = true
	config.Producer.Return.Successes = true
	config.Net.WriteTimeout = 5 * time.Second
	config.Producer.Retry.Max = 0

	if t.Username != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = t.Username
		config.Net.SASL.Password = t.Password
	}
	return config

}

// Consumer ...
func (t *KafkaHelper) Consumer() {
	// brokers := []string{"kafka-1:9092", "kafka-2:9092"}
	// group := "Your-Consumer-Group"
	// topics := []string{"topicName"}
	// consumer := cluster.NewConsumer(broker, group, topics, conf)

	// Create new consumer topic
	master, err := sarama.NewConsumer([]string{t.Host}, t.config())
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	topics := []string{"users"}
	// topics, _ := master.Topics(topic, nil)
	consumer, errors := consume(topics, master)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Count how many message processed
	msgCount := 0

	// Get signnal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case msg := <-consumer:
				msgCount++
				fmt.Println(fmt.Sprintf("Offset:%v Partition:%v Key:%v Topic:%v Time:%v", msg.Offset, msg.Partition, string(msg.Key), msg.Topic, msg.Timestamp))
				fmt.Println(fmt.Sprintf("Messages: %v", string(msg.Value)))
			case consumerError := <-errors:
				msgCount++
				fmt.Println("Received consumerError ", string(consumerError.Topic), string(consumerError.Partition), consumerError.Err)
				doneCh <- struct{}{}
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

}

func consume(topics []string, master sarama.Consumer) (chan *sarama.ConsumerMessage, chan *sarama.ConsumerError) {
	consumers := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)
	for _, topic := range topics {
		if strings.Contains(topic, "__consumer_offsets") {
			continue
		}
		partitions, _ := master.Partitions(topic)
		// this only consumes partition no 1, you would probably want to consume all partitions
		consumer, err := master.ConsumePartition(topic, partitions[0], sarama.OffsetOldest)
		if nil != err {
			fmt.Printf("Topic %v Partitions: %v", topic, partitions)
			panic(err)
		}
		fmt.Println("Start consuming topic ", topic)
		go func(topic string, consumer sarama.PartitionConsumer) {
			for {
				select {
				case consumerError := <-consumer.Errors():
					errors <- consumerError
				case msg := <-consumer.Messages():
					consumers <- msg
				}
			}
		}(topic, consumer)
	}

	return consumers, errors
}
