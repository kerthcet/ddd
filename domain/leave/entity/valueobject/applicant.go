package valueobject

import (
	personENT "ddd/domain/person/entity"
)

// Applicant 申请人
// 值对象没有唯一标识，但是不妨碍它的属性是一个实体
type Applicant struct {
	PersonID   int
	PersonName string
	PersonType int
}

// ToApplicant 转化为applicant
func ToApplicant(p *personENT.Person) *Applicant {
	return &Applicant{
		PersonID:   p.ID,
		PersonName: p.Name,
	}
}
