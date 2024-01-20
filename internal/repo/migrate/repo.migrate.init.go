package migrate

import (
	"github.com/huseinnashr/pforder-backend/internal/config"
	"github.com/huseinnashr/pforder-backend/internal/domain"
)

type Repo struct {
	config *config.Config
}

func New(config *config.Config) domain.IMigrateRepo {
	return &Repo{
		config: config,
	}

}
