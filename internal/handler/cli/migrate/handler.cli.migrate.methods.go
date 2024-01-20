package migrate

import (
	"github.com/huseinnashr/pforder-backend/internal/pkg/tracer"
	"github.com/spf13/cobra"
)

func (h *Handler) MigrateAll(cmd *cobra.Command, args []string) error {
	ctx, span := tracer.Start(cmd.Context(), "handler.MigrateAll")
	defer span.End()

	return h.migrateUsecase.MigrateAll(ctx)
}
