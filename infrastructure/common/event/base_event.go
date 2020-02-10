package event

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
)

// DomainEvent 事件基本字段
type DomainEvent struct {
	ID              string
	CreatedAt       time.Time
	Metadata        interface{}
	ServiceName     string
	AggregationName string
	Action          string
}

// NewEvent 返回基础事件
func NewEvent() *DomainEvent {
	// TODO servicename从env获取
	return &DomainEvent{
		ID:          "123456789",
		ServiceName: "leave",
		CreatedAt:   time.Now(),
	}
}

// ToBytes 转化
func (e *DomainEvent) ToBytes() ([]byte, error) {
	res, err := json.Marshal(e)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return res, nil
}
