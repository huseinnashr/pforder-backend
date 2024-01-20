package main

import (
	"context"

	"github.com/huseinnashr/pforder-backend/internal/config"
	"github.com/huseinnashr/pforder-backend/internal/pkg/tracer"
)

const APP_NAME = "bimble-backend-http"

func main() {
	var ctx = context.Background()

	config, err := config.GetConfig("./files/config/development.yaml")
	if err != nil {
		panic(err)
	}

	tracerShutdown, err := tracer.Init(ctx, config, APP_NAME)
	if err != nil {
		panic(err)
	}
	defer tracerShutdown(ctx)

	if err := startApp(ctx, config); err != nil {
		panic(err)
	}
}
