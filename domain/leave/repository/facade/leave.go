package facade

import (
	"github.com/golobby/container"

	leavePST "ddd/domain/leave/repository/persistence/mysql"
	leavePO "ddd/domain/leave/repository/po"
)

// Repo 仓储
var Repo LeaveRepository

func init() {
	container.Singleton(func() LeaveRepository {
		return &leavePST.RepoImpl{}
	})
	container.Make(&Repo)
}

// LeaveRepository 仓储接口
type LeaveRepository interface {
	Create(*leavePO.Leave) (*leavePO.Leave, error)
	CreateEvent(*leavePO.LeaveEvent) (*leavePO.LeaveEvent, error)
}
