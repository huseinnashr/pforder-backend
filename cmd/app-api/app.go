package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huseinnashr/pforder-backend/internal/config"
	orderhandler "github.com/huseinnashr/pforder-backend/internal/handler/api/order"
	orderrepo "github.com/huseinnashr/pforder-backend/internal/repo/order"
	orderusecase "github.com/huseinnashr/pforder-backend/internal/usecase/order"
	_ "github.com/lib/pq"
)

func startApp(ctx context.Context, config *config.Config) error {
	sqlDSN := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Resource.SQLDatabase.Host, config.Resource.SQLDatabase.Port, config.Resource.SQLDatabase.User,
		config.Resource.SQLDatabase.Password, config.Resource.SQLDatabase.DBName,
	)
	sqlDatabase, err := sql.Open("postgres", sqlDSN)
	if err != nil {
		return err
	}

	orderRepo := orderrepo.New(config, sqlDatabase)
	orderUsecase := orderusecase.New(orderRepo)
	orderHandler := orderhandler.New(orderUsecase)

	if err := startServer(ctx, config, orderHandler); err != nil {
		return err
	}

	return nil
}
