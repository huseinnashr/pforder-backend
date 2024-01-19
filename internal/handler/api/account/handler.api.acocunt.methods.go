package account

import (
	"context"
	"log"

	v1 "github.com/huseinnashr/pforder-backend/api/v1"
	"github.com/huseinnashr/pforder-backend/internal/domain"
	"github.com/huseinnashr/pforder-backend/internal/pkg/tracer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Signup implements v1.AccountServiceHTTPServer.
func (h *Handler) ListOrder(ctx context.Context, req *v1.ListOrderRequest) (*v1.ListOrderResponse, error) {
	ctx, span := tracer.Start(ctx, "handler.Signup")
	defer span.End()

	log.Println(req.GetFilters().GetEndDate().AsTime())
	orders, nextCursor, err := h.orderUsecase.ListOrder(ctx, domain.ListOrderParam{
		Search:    req.GetFilters().GetSearch(),
		StartDate: req.GetFilters().GetStartDate().AsTime(),
		EndDate:   req.GetFilters().GetEndDate().AsTime(),
		PageSize:  req.GetPagination().GetSize(),
		Cursor:    req.GetPagination().GetCursor(),
	})
	if err != nil {
		return nil, err
	}

	pbOrders := make([]*v1.Order, len(orders))
	for i, order := range orders {
		pbOrders[i] = &v1.Order{
			OrderName:           order.OrderName,
			CustomerCompanyName: order.CustomerCompanyName,
			OrderDate:           timestamppb.New(order.OrderDate),
			CustomerName:        order.CustomerName,
			DeliveredAmount:     order.DeliveredAmount,
			TotalAmount:         order.TotalAmount,
		}
	}

	return &v1.ListOrderResponse{
		Orders: pbOrders,
		Pagination: &v1.PaginationResponse{
			NextCursor: nextCursor,
		},
	}, nil
}
