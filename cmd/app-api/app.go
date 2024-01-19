package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huseinnashr/pforder-backend/internal/config"
	orderhandler "github.com/huseinnashr/pforder-backend/internal/handler/api/account"
	orderrepo "github.com/huseinnashr/pforder-backend/internal/repo/account"
	orderusecase "github.com/huseinnashr/pforder-backend/internal/usecase/account"
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
