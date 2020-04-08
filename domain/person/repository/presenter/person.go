package presenter

import (
	"github.com/golobby/container"

	personPST "ddd/domain/person/repository/persistence/mysql"
	personPO "ddd/domain/person/repository/po"
)

// Repo 仓储
var Repo PersonRepository

func init() {
	container.Singleton(func() PersonRepository {
		return &personPST.RepoImpl{}
	})
	container.Make(&Repo)
}

// PersonRepository 仓储接口
type PersonRepository interface {
	FindPersonByID(int) (*personPO.Person, error)
}
