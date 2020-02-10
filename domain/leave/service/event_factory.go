package service

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	leavePO "ddd/domain/leave/repository/po"
	commonEvent "ddd/infrastructure/common/event"
	commonPO "ddd/infrastructure/common/po"
)

const (
	// AggregationName 聚合名称
	AggregationName = "leave"
)

// NewLeaveEventDO 新建leaveEventDO
func NewLeaveEventDO(leave *leavePO.Leave, action string) *commonEvent.DomainEvent {
	event := commonEvent.NewEvent()
	event.AggregationName = AggregationName
	event.Action = action
	event.Metadata = leave

	return event
}

// ToLeaveEventPO DO -> PO
func ToLeaveEventPO(e *commonEvent.DomainEvent) *leavePO.LeaveEvent {
	res, err := json.Marshal(e)
	if err != nil {
		log.Error(err)
		return nil
	}

	return &leavePO.LeaveEvent{
		Event: commonPO.Event{
			ID:            1,
			CreatedAt:     e.CreatedAt,
			DomainEventID: e.ID,
			Metadata:      string(res),
		},
	}
}
