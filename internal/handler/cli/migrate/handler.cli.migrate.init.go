package migrate

import (
	"github.com/huseinnashr/pforder-backend/internal/domain"
	"github.com/spf13/cobra"
)

type Handler struct {
	migrateUsecase domain.IMigrateUsecase
}

func New(migrateUsecase domain.IMigrateUsecase) *cobra.Command {
	handler := &Handler{
		migrateUsecase: migrateUsecase,
	}

	command := &cobra.Command{Use: "migrate"}
	command.AddCommand(&cobra.Command{
		Use:  "all",
		RunE: handler.MigrateAll,
	})

	return command
}
