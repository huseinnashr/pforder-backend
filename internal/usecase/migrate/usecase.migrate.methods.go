package migrate

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/huseinnashr/pforder-backend/internal/domain"
	"github.com/huseinnashr/pforder-backend/internal/pkg/tracer"
)

func (u *Usecase) MigrateAll(ctx context.Context) error {
	ctx, span := tracer.Start(ctx, "usecase.MigrateAll")
	defer span.End()

	dbTx, err := u.sqlDatabase.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if err := dbTx.Rollback(); err != nil {
				log.Println("error rollback", err)
			}
		} else {
			if err := dbTx.Commit(); err != nil {
				log.Println("error commit", err)
			}
		}
	}()

	transformArray := func(value string) (interface{}, error) {
		replacer := strings.NewReplacer("[", "{", "]", "}")
		return replacer.Replace(value), nil
	}

	transformMoney := func(value string) (interface{}, error) {
		if value == "" {
			return int64(0), nil
		}

		floatMoney, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		return int64(floatMoney * float64(domain.SMALLEST_MONEY_UNIT)), nil
	}

	migrateParams := []domain.CSVToDBParam{
		{
			Filepath: "./files/data/customer_companies.csv", TableName: "customer_companies",
			Columns: []string{"company_id", "company_name"},
		},
		{
			Filepath: "./files/data/customers.csv", TableName: "customers",
			Columns:   []string{"user_id", "login", "password", "name", "company_id", "credit_cards"},
			Transform: map[string]domain.Transformer{"credit_cards": transformArray},
		},
		{
			Filepath: "./files/data/orders.csv", TableName: "orders",
			Columns: []string{"id", "created_at", "order_name", "customer_id"},
		},
		{
			Filepath: "./files/data/order_items.csv", TableName: "order_items",
			Columns:   []string{"id", "order_id", "price_per_unit", "quantity", "product"},
			Transform: map[string]domain.Transformer{"price_per_unit": transformMoney},
		},
		{
			Filepath: "./files/data/deliveries.csv", TableName: "deliveries",
			Columns: []string{"id", "order_item_id", "delivered_quantity"},
		},
	}

	for _, migrateParam := range migrateParams {
		if err = u.migrateRepo.CSVToDB(ctx, dbTx, migrateParam); err != nil {
			return err
		}
	}

	log.Println("Migrate Success")
	return nil
}
