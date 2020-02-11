package migrate

import (
	log "github.com/sirupsen/logrus"

	leavePO "ddd/domain/leave/repository/po"
	personPO "ddd/domain/person/repository/po"
	"ddd/infrastructure/util/driver"
)

// Run migrate
func Run() {
	log.Info("Starting to migrate")
	driver.DB.AutoMigrate(&personPO.Person{})
	driver.DB.AutoMigrate(&leavePO.Leave{})
	driver.DB.AutoMigrate(&leavePO.LeaveEvent{})
}
