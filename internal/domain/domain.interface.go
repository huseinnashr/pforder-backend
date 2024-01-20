package domain

import (
	"context"
	"database/sql"
	"time"
)

const SMALLEST_MONEY_UNIT int64 = 1000

type IOrderUsecase interface {
	ListOrder(ctx context.Context, params ListOrderParam) ([]Order, string, error)
}

type IMigrateUsecase interface {
	MigrateAll(ctx context.Context) error
}

type IOrderRepo interface {
	ListOrder(ctx context.Context, params ListOrderParam) ([]Order, string, error)
}

type IMigrateRepo interface {
	CSVToDB(ctx context.Context, dbTx ISQLDatabaseTx, params CSVToDBParam) error
}

type ISQLDatabase interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}
type ISQLDatabaseTx interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Commit() error
	Rollback() error
}

type ListOrderParam struct {
	Search    string
	StartDate time.Time
	EndDate   time.Time
	PageSize  int64
	Cursor    string
}

type Transformer func(value string) (interface{}, error)
type CSVToDBParam struct {
	Filepath  string
	TableName string
	Columns   []string
	Transform map[string]Transformer
}

type Order struct {
	OrderName           string
	CustomerCompanyName string
	CustomerName        string
	OrderDate           time.Time
	DeliveredAmount     int64
	TotalAmount         int64
}
