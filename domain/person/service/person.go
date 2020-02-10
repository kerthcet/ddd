package service

import (
	log "github.com/sirupsen/logrus"

	personENT "ddd/domain/person/entity"
	personRepo "ddd/domain/person/repository/facade"
)

// FindPersonByID 根据id获得
func FindPersonByID(id int) (*personENT.Person, error) {
	personPO, err := personRepo.Repo.FindPersonByID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return ToPersonDO(personPO), nil
}
