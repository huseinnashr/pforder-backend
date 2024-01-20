package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huseinnashr/pforder-backend/internal/config"
	migratehandler "github.com/huseinnashr/pforder-backend/internal/handler/cli/migrate"
	migraterepo "github.com/huseinnashr/pforder-backend/internal/repo/migrate"
	migrateusecase "github.com/huseinnashr/pforder-backend/internal/usecase/migrate"
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
	migrateRepo := migraterepo.New(config)
	migrateUsecase := migrateusecase.New(migrateRepo, sqlDatabase)
	migrateHandler := migratehandler.New(migrateUsecase)

	if err := startCommand(ctx, migrateHandler); err != nil {
		return err
	}

	return nil
}
