package driver

import (
	"context"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	log "github.com/sirupsen/logrus"

	"ddd/infrastructure/config"
)

var produceTopics = []string{}

// producer 生产者
var producer *kafkaProducer

func init() {
	producer = &kafkaProducer{
		topics:  produceTopics,
		writers: make(map[string]*kafka.Writer),
	}

	for i := range producer.topics {
		topic := producer.topics[i]
		writer := kafka.NewWriter(newKafkaWriterConfig(topic))
		producer.writers[topic] = writer
	}

	log.Infof("Produce to topics: %v", produceTopics)
}

// kafkaProducer 生产者
type kafkaProducer struct {
	topics  []string
	writers map[string]*kafka.Writer
}

func newKafkaWriterConfig(topic string) kafka.WriterConfig {
	return kafka.WriterConfig{
		Brokers:           []string{config.ENV.KafkaHost},
		Topic:             topic,
		BatchSize:         100,
		BatchTimeout:      100 * time.Millisecond,
		BatchBytes:        1048576 * 10, // 10M
		Async:             false,
		RequiredAcks:      1,
		RebalanceInterval: 0,
		MaxAttempts:       5,
		CompressionCodec:  snappy.NewCompressionCodec(),
	}
}

// Produce 控制生产者个数
func Produce(topic string, eventBytes []byte) (err error) {
	log.Infof("write msg to kafka topic `%s`: %s ", topic, string(eventBytes))

	writer := producer.writers[topic]
	err = writer.WriteMessages(context.Background(),
		kafka.Message{Value: eventBytes},
	)
	if err != nil {
		log.Error("Write msg to kafka error: ", err)
		return err
	}

	return nil
}

// CloseAll 关闭所有生产者
// TODO 优雅的关闭
func CloseAll() {
	for i := range producer.topics {
		topic := producer.topics[i]
		producer.writers[topic].Close()
	}
}
