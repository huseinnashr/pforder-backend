package account

import (
	v1 "github.com/huseinnashr/pforder-backend/api/v1"
	"github.com/huseinnashr/pforder-backend/internal/domain"
)

type Handler struct {
	v1.UnimplementedAccountServiceServer
	orderUsecase domain.IOrderUsecase
}

func New(orderUsecase domain.IOrderUsecase) v1.AccountServiceServer {
	return &Handler{
		orderUsecase: orderUsecase,
	}
}
