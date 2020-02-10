package driver

import (
	"context"
	"fmt"
	"sync"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	log "github.com/sirupsen/logrus"

	"ddd/infrastructure/common/event"
	"ddd/infrastructure/config"
)

// producer 生产者
var producer *kafkaProducer
var mutex sync.RWMutex

func init() {
	producer = &kafkaProducer{}
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
		BatchBytes:        1048576 * 10, // 10M
		BatchTimeout:      5 * time.Second,
		RequiredAcks:      1,
		Async:             false,
		RebalanceInterval: 300 * time.Second,
		CompressionCodec:  snappy.NewCompressionCodec(),
	}
}

// Produce 控制生产者个数
func Produce(topic string, event *event.DomainEvent) error {
	writer, err := producer.findOrCreateWriter(topic)
	if err != nil {
		return err
	}

	eventBytes, err := event.ToBytes()
	if err != nil {
		return err
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{Value: eventBytes},
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *kafkaProducer) findOrCreateWriter(topic string) (*kafka.Writer, error) {
	var err error
	if !p.hasMutexWriter(topic) {
		err = p.addWriter(topic)
	}

	if err != nil {
		return nil, err
	}
	return p.writers[topic], nil
}

func (p *kafkaProducer) hasMutexWriter(topic string) bool {
	defer func() {
		mutex.RUnlock()
	}()

	mutex.RLock()
	return p.hasWriter(topic)
}

func (p *kafkaProducer) hasWriter(topic string) bool {
	if p.writers[topic] == nil {
		return false
	}

	return true
}

func (p *kafkaProducer) addWriter(topic string) (err error) {
	defer func() {
		mutex.Unlock()
		if r := recover(); r != nil {
			log.Error(r)
			err = fmt.Errorf("New kafka writer error: %s", r)
			return
		}
	}()

	mutex.Lock()
	// 二次检查
	if p.hasWriter(topic) {
		return
	}

	writer := kafka.NewWriter(newKafkaWriterConfig(topic))
	p.writers[topic] = writer
	p.topics = append(p.topics, topic)
	return
}

// CloseAll 关闭所有生产者
// TODO 优雅的关闭
func CloseAll() {
	for i := range producer.topics {
		topic := producer.topics[i]
		producer.writers[topic].Close()
	}
}
