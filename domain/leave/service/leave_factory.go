package service

import (
	leaveEntity "ddd/domain/leave/entity"
	leaveVO "ddd/domain/leave/entity/valueobject"
	leavePO "ddd/domain/leave/repository/po"
)

// ToLeavePO DO -> PO
func ToLeavePO(leave *leaveEntity.LeaveDO) *leavePO.Leave {
	return &leavePO.Leave{
		ApplicantID:   leave.Applicant.PersonID,
		ApplicantName: leave.Applicant.PersonName,
		ApproverID:    leave.Approver.PersonID,
		ApproverName:  leave.Approver.PersonName,
		ApprovalType:  leaveVO.APPROVING,
		StartTime:     leave.StartTime,
		EndTime:       leave.EndTime,
	}
}

// ToLeaveDO PO -> DO
func ToLeaveDO(leave *leavePO.Leave, approver *leaveVO.Approver, applicant *leaveVO.Applicant) *leaveEntity.LeaveDO {
	return &leaveEntity.LeaveDO{
		Approver:     approver,
		Applicant:    applicant,
		ApprovalType: leave.ApprovalType,
		StartTime:    leave.StartTime,
		EndTime:      leave.EndTime,
	}
}
