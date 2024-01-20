package order

import (
	"github.com/huseinnashr/pforder-backend/internal/config"
	"github.com/huseinnashr/pforder-backend/internal/domain"
)

type Repo struct {
	config      *config.Config
	sqlDatabase domain.ISQLDatabase
}

func New(config *config.Config, sqlDatabase domain.ISQLDatabase) domain.IOrderRepo {
	return &Repo{
		config:      config,
		sqlDatabase: sqlDatabase,
	}

}
