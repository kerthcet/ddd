package service

import (
	personENT "ddd/domain/person/entity"
	personRepo "ddd/domain/person/repository/presenter"
)

// FindPersonByID 根据id获得
func FindPersonByID(id int) (*personENT.Person, error) {
	personPO, err := personRepo.Repo.FindPersonByID(id)
	if err != nil {
		return nil, err
	}
	return ToPersonDO(personPO), nil
}
