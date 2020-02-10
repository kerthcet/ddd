package po

import (
	"time"

	"ddd/infrastructure/common/po"
)

// Leave 请假单po
type Leave struct {
	ID            int
	ApplicantID   int
	ApplicantName string
	ApproverID    int
	ApproverName  string
	ApprovalType  int
	StartTime     time.Time
	EndTime       time.Time
}

// LeaveEvent 请假事件po
type LeaveEvent struct {
	po.Event
}
