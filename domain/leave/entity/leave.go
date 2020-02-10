package entity

import (
	leaveVO "ddd/domain/leave/entity/valueobject"
	"time"
)

// LeaveDO 请假聚合根
type LeaveDO struct {
	ID           int
	Approver     *leaveVO.Approver
	Applicant    *leaveVO.Applicant
	ApprovalType int
	StartTime    time.Time
	EndTime      time.Time
}

// GetDuration 请假时长
func (l *LeaveDO) GetDuration() time.Duration {
	return l.EndTime.Sub(l.StartTime)
}

// SetApprover 设置approver
func (l *LeaveDO) SetApprover(leave *leaveVO.Approver) {
	l.Approver = leave
}

// SetApplicant 设置applicant
func (l *LeaveDO) SetApplicant(leave *leaveVO.Applicant) {
	l.Applicant = leave
}
