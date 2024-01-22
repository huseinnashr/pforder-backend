package order

import (
	"testing"

	"github.com/huseinnashr/pforder-backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type handlerMock struct {
	orderUsecase *domain.MockIOrderUsecase
}

func newHandlerMock(ctrl *gomock.Controller) *handlerMock {
	return &handlerMock{
		orderUsecase: domain.NewMockIOrderUsecase(ctrl),
	}
}

func (h *handlerMock) toHandler() *Handler {
	return &Handler{
		orderUsecase: h.orderUsecase,
	}
}

func TestNew(t *testing.T) {
	got := New(nil)
	assert.NotNil(t, got, "it should return handler")
}
