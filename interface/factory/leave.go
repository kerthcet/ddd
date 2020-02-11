package factory

import (
	leaveEntity "ddd/domain/leave/entity"
	leaveVO "ddd/domain/leave/entity/valueobject"
	"ddd/interface/dto/rest"
)

// ToCreateLeaveDO DTO -> DO
func ToCreateLeaveDO(leave *rest.CreateLeaveRequestDTO) *leaveEntity.LeaveDO {
	applicant := &leaveVO.Applicant{
		PersonID: leave.ApplicantID,
	}

	approver := &leaveVO.Approver{
		PersonID: leave.ApproverID,
	}

	return &leaveEntity.LeaveDO{
		Applicant:    applicant,
		Approver:     approver,
		ApprovalType: leave.ApprovalType,
		StartTime:    leave.StartTime,
		EndTime:      leave.EndTime,
	}
}

// ToCreateLeaveDTO DO -> DTO
func ToCreateLeaveDTO(leave *leaveEntity.LeaveDO) *rest.CreateLeaveResponseDTO {
	return &rest.CreateLeaveResponseDTO{
		LeaveID:       leave.ID,
		StartTime:     leave.StartTime,
		EndTime:       leave.EndTime,
		ApplicantID:   leave.Applicant.PersonID,
		ApplicantName: leave.Applicant.PersonName,
		ApproverID:    leave.Approver.PersonID,
		ApproverName:  leave.Approver.PersonName,
	}
}
