package entity

import (
	"time"

	leaveVO "ddd/domain/leave/entity/valueobject"
)

// Approval 审批意见实体
type Approval struct {
	Approver     *leaveVO.Approver
	ApprovalID   int
	ApprovalType int
	Msg          string
	Durationn    *time.Duration
}
