package driver

import (
	"context"
	"sync"
	"time"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"

	"ddd/infrastructure/config"
)

var consumerTopics = []string{}

var consumer *kafkaConsumer
var wg sync.WaitGroup

const (
	consumerGroup = "ddd-group"
)

// kafkaConsumer 消费者
type kafkaConsumer struct {
	topics  []string
	readers map[string]*kafka.Reader
}

func init() {
	consumer = &kafkaConsumer{
		topics:  consumerTopics,
		readers: make(map[string]*kafka.Reader),
	}

	for i := range consumer.topics {
		topic := consumer.topics[i]
		consumer.readers[topic] = newConsumer(topic)
	}

	log.Infof("Consume topics: %v", consumerTopics)
}

// Consume 消费kafka消息
func Consume() {
	consumer.run()
}

func (kc *kafkaConsumer) run() {
	for i := range kc.readers {
		wg.Add(1)
		go run(kc.readers[i])
	}

	wg.Wait()
}

func run(r *kafka.Reader) {
	defer wg.Done()
	defer r.Close()

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			log.Error("Read msg error: ", err)
			continue
		}
		log.Infof("read message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		r.CommitMessages(ctx, m)
	}
}

func newConsumer(topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{config.ENV.KafkaHost},
		GroupID:        consumerGroup,
		Topic:          topic,
		MinBytes:       10e3,            // 10KB
		MaxBytes:       10e6,            // 10MB
		CommitInterval: 3 * time.Second, // autocommit
	})
}
