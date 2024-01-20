package migrate

import (
	"github.com/huseinnashr/pforder-backend/internal/domain"
)

type Usecase struct {
	migrateRepo domain.IMigrateRepo
	sqlDatabase domain.ISQLDatabase
}

func New(migrateRepo domain.IMigrateRepo, sqlDatabase domain.ISQLDatabase) domain.IMigrateUsecase {
	return &Usecase{
		migrateRepo: migrateRepo,
		sqlDatabase: sqlDatabase,
	}
}
