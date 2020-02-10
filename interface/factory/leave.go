package factory

import (
	leaveEntity "ddd/domain/leave/entity"
	leaveVO "ddd/domain/leave/entity/valueobject"
	"ddd/interface/dto"
)

// ToCreateLeaveDO DTO -> DO
func ToCreateLeaveDO(leave *dto.CreateLeaveRequestDTO) *leaveEntity.LeaveDO {
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
func ToCreateLeaveDTO(leave *leaveEntity.LeaveDO) *dto.CreateLeaveResponseDTO {
	return &dto.CreateLeaveResponseDTO{
		LeaveID:       leave.ID,
		StartTime:     leave.StartTime,
		EndTime:       leave.EndTime,
		ApplicantID:   leave.Applicant.PersonID,
		ApplicantName: leave.Applicant.PersonName,
		ApproverID:    leave.Approver.PersonID,
		ApproverName:  leave.Approver.PersonName,
	}
}
