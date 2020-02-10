package mysql

import (
	log "github.com/sirupsen/logrus"

	leavePO "ddd/domain/leave/repository/po"
	"ddd/infrastructure/util/driver"
)

// RepoImpl 仓储实现
type RepoImpl struct{}

// Create 创建leave
func (r *RepoImpl) Create(leave *leavePO.Leave) (*leavePO.Leave, error) {
	err := driver.DB.Create(leave).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return leave, nil
}

// CreateEvent 创建leave event
func (r *RepoImpl) CreateEvent(event *leavePO.LeaveEvent) (*leavePO.LeaveEvent, error) {
	err := driver.DB.Create(event).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return event, nil
}
