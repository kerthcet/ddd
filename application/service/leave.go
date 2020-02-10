package service

import (
	leaveENT "ddd/domain/leave/entity"
	// personVO "ddd/domain/leave/entity/valueobject"
	// leaveDS "ddd/domain/leave/service"
	personDS "ddd/domain/person/service"
)

// CreateLeave 创建leave
func CreateLeave(leave *leaveENT.LeaveDO) (*leaveENT.LeaveDO, error) {
	_, err := personDS.FindPersonByID(leave.Approver.PersonID)
	if err != nil {
		return nil, err
	}

	// applicant, err := personDS.FindPersonByID(leave.Applicant.PersonID)
	// if err != nil {
	//     return nil, err
	// }

	// leave.SetApprover(personVO.ToApprover(approver))
	// leave.SetApplicant(personVO.ToApplicant(applicant))
	// leave, err = leaveDS.CreateLeave(leave)
	// if err != nil {
	//     return nil, err
	// }

	return leave, nil
}
