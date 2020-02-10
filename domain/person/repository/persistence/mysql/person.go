package mysql

import (
	log "github.com/sirupsen/logrus"

	personPO "ddd/domain/person/repository/po"
	"ddd/infrastructure/util/driver"
)

// RepoImpl 仓储实现
type RepoImpl struct{}

// FindPersonByID 根据id找到对应的person
func (r *RepoImpl) FindPersonByID(id int) (*personPO.Person, error) {
	var person personPO.Person
	log.Error("jjjjjjjjjjjjjjjjjjj", driver.DB)
	err := driver.DB.Model(&personPO.Person{}).
		Where("id = ?", id).
		First(&person).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &person, nil
}
