package valueobject

import (
	personENT "ddd/domain/person/entity"
)

// Approver 审批人
type Approver struct {
	PersonID   int
	PersonName string
}

// ToApprover 转化为approver
func ToApprover(p *personENT.Person) *Approver {
	return &Approver{
		PersonID:   p.ID,
		PersonName: p.Name,
	}
}
