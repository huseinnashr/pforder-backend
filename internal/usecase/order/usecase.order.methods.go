package order

import (
	"context"

	"github.com/huseinnashr/pforder-backend/internal/domain"
	"github.com/huseinnashr/pforder-backend/internal/pkg/tracer"
)

func (u *Usecase) ListOrder(ctx context.Context, params domain.ListOrderParam) ([]domain.Order, string, error) {
	ctx, span := tracer.Start(ctx, "usecase.ListOrder")
	defer span.End()

	return u.orderRepo.ListOrder(ctx, params)
}
