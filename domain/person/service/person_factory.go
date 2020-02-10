package service

import (
	"time"

	personEntity "ddd/domain/person/entity"
	personPO "ddd/domain/person/repository/po"
)

// ToPersonPO DO -> PO
func ToPersonPO(p *personEntity.Person) *personPO.Person {
	personPO := &personPO.Person{}
	personPO.ID = p.ID
	personPO.Name = p.Name
	personPO.CreatedAt = time.Now()
	personPO.UpdatedAt = time.Now()

	return personPO
}

// ToPersonDO PO -> DO
func ToPersonDO(p *personPO.Person) *personEntity.Person {
	return &personEntity.Person{
		ID:   p.ID,
		Name: p.Name,
	}
}
