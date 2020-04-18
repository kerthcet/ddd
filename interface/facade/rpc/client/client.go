package client

import (
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"ddd/interface/dto/rpc"
)

const (
	address = "localhost:9090"
)

// Send 发送
func Send() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Duration(10*time.Second)))
	if err != nil {
		log.Error(err)
		return
	}
	defer conn.Close()

	client := rpc.NewLeaverClient(conn)

	resp, err := client.CreateLeave(context.Background(), &rpc.CreateLeaveRequest{
		ApplicantID:  1,
		ApproverID:   1,
		ApprovalType: 1,
		StartTime:    "2018-08-13 12:23:34",
		EndTime:      "2019-11-13 22:23:34",
	})
	if err != nil {
		log.Error(err)
	}
	log.Println(resp.GetLeaveID())
}
