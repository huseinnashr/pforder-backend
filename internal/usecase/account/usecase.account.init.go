package account

import (
	"github.com/huseinnashr/pforder-backend/internal/domain"
)

type Usecase struct {
	orderRepo domain.IOrderRepo
}

func New(accountRepo domain.IOrderRepo) domain.IOrderUsecase {
	return &Usecase{
		orderRepo: accountRepo,
	}
}
