package order

import (
	"context"

	"github.com/huseinnashr/pforder-backend/internal/domain"
	"github.com/huseinnashr/pforder-backend/internal/pkg/tracer"
)

func (r *Repo) ListOrder(ctx context.Context, params domain.ListOrderParam) ([]domain.Order, string, error) {
	ctx, span := tracer.Start(ctx, "repo.ListOrder")
	defer span.End()

	row := r.sqlDatabase.QueryRowContext(ctx,
		``,
	)
	if err := row.Err(); err != nil {
		return nil, "", err
	}

	var accountID int64
	if err := row.Scan(&accountID); err != nil {
		return nil, "", err
	}

	return nil, "", nil
}
