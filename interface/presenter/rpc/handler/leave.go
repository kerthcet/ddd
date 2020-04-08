package handler

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"ddd/interface/dto/rpc"
)

// CreateLeave 创建leave
type CreateLeave struct{}

// CreateLeave 创建请假单
func (c *CreateLeave) CreateLeave(ctx context.Context, in *rpc.CreateLeaveRequest) (*rpc.CreateLeaveResponse, error) {
	log.Infof("Received: %+v", in)
	out := &rpc.CreateLeaveResponse{}
	return out, nil
}
