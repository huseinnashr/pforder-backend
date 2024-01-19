package domain

import (
	"context"
	"database/sql"
	"time"
)

type IOrderUsecase interface {
	ListOrder(ctx context.Context, params ListOrderParam) ([]Order, string, error)
}

type IOrderRepo interface {
	ListOrder(ctx context.Context, params ListOrderParam) ([]Order, string, error)
}

type ISQLDatabase interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type ListOrderParam struct {
	Search    string
	StartDate time.Time
	EndDate   time.Time
	PageSize  int64
	Cursor    string
}

type Order struct {
	OrderName           string
	CustomerCompanyName string
	CustomerName        string
	OrderDate           time.Time
	DeliveredAmount     int64
	TotalAmount         int64
}
