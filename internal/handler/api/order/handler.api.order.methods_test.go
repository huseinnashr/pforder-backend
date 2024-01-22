package order

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	v1 "github.com/huseinnashr/pforder-backend/api/v1"
	"github.com/huseinnashr/pforder-backend/internal/config"
	"github.com/huseinnashr/pforder-backend/internal/domain"
	"github.com/huseinnashr/pforder-backend/internal/pkg/tracer"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestMain(m *testing.M) {
	tracer.Init(context.Background(), &config.Config{}, "")
	os.Exit(m.Run())
}

func TestListOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type TestCase struct {
		name    string
		params  *v1.ListOrderRequest
		mockFn  func(h *handlerMock)
		want    *v1.ListOrderResponse
		wantErr bool
	}

	testCases := []TestCase{
		{
			name:   "Param nil, listOrder return err",
			params: nil,
			mockFn: func(h *handlerMock) {
				h.orderUsecase.EXPECT().ListOrder(gomock.Any(), gomock.Any()).Return(nil, "", errors.New("error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Success",
			params: &v1.ListOrderRequest{Filters: &v1.ListOrderRequest_Filter{Search: "A"}},
			mockFn: func(h *handlerMock) {
				h.orderUsecase.EXPECT().
					ListOrder(gomock.Any(), domain.ListOrderParam{Search: "A", StartDate: time.Unix(0, 0).UTC(), EndDate: time.Unix(0, 0).UTC()}).
					Return([]domain.Order{{OrderName: "order_name", OrderDate: time.Unix(0, 0).UTC()}}, "cursor", nil)
			},
			want: &v1.ListOrderResponse{
				Orders:     []*v1.Order{{OrderName: "order_name", OrderDate: timestamppb.New(time.Unix(0, 0).UTC())}},
				Pagination: &v1.PaginationResponse{NextCursor: "cursor"},
			},
		},
	}

	for _, tt := range testCases {
		handlerMock := newHandlerMock(ctrl)
		tt.mockFn(handlerMock)

		handler := handlerMock.toHandler()
		got, err := handler.ListOrder(context.Background(), tt.params)
		if tt.wantErr != (err != nil) {
			t.Errorf("Expect err %v got %v", tt.wantErr, err.Error())
		}
		assert.Equal(t, tt.want, got, "it should be equal")
	}
}
