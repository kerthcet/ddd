package service

import (
	leaveEntity "ddd/domain/leave/entity"
	leaveEvent "ddd/domain/leave/event"
	leaveRepo "ddd/domain/leave/repository/presenter"
	"ddd/infrastructure/util/driver"
)

// CreateLeave 创建leave
func CreateLeave(leave *leaveEntity.LeaveDO) (*leaveEntity.LeaveDO, error) {
	// TODO 事务
	// 持久化业务实体
	leavePO, err := leaveRepo.Repo.Create(ToLeavePO(leave))
	if err != nil {
		return nil, err
	}
	// 持久化事件实体
	eventDO := NewLeaveEventDO(leavePO, leaveEvent.CreateLeaveEvent)
	_, err = leaveRepo.Repo.CreateEvent(ToLeaveEventPO(eventDO))
	if err != nil {
		return nil, err
	}

	// 发送事件
	err = driver.Produce(leaveEvent.EventTopicName, eventDO)
	if err != nil {
		return nil, err
	}
	return ToLeaveDO(leavePO, leave.Approver, leave.Applicant), nil
}
