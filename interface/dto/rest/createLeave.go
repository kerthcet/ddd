package rest

import (
	"time"
)

// CreateLeaveRequestDTO DTO
type CreateLeaveRequestDTO struct {
	ApplicantID  int       `json:"applicantId" binding:"required"`
	ApproverID   int       `json:"approverId" binding:"required"`
	ApprovalType int       `json:"approvalType" binding:"required"`
	StartTime    time.Time `json:"startTime" binding:"required" time_format:"2006-01-02T13:20:34Z"`
	EndTime      time.Time `json:"endTime" binding:"required" time_format:"2006-01-02T13:20:34Z"`
}

// CreateLeaveResponseDTO DTO
type CreateLeaveResponseDTO struct {
	LeaveID       int       `json:"leaveId"`
	ApplicantID   int       `json:"applicanntId"`
	ApplicantName string    `json:"applicantName"`
	ApproverID    int       `json:"approverId"`
	ApproverName  string    `jsonn:"approverName"`
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
}
