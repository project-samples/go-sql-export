package main

import (
	"context"

	"github.com/core-go/config"
	"github.com/core-go/log/rotatelogs"
	log "github.com/core-go/log/zap"

	"go-service/internal/app"
)

func main() {
	var cfg app.Config
	err := config.Load(&cfg, "configs/config")
	if err != nil {
		panic(err)
	}
	_, err = log.InitializeWithWriter(cfg.Log, rotatelogs.GetWriter)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	log.Info(ctx, "Export file")
	app, err := app.NewApp(ctx, cfg)
	if err != nil {
		log.Errorf(ctx, "Error when initialize: %v", err)
		panic(err)
	}
	total, err := app.Export(ctx)
	if err != nil {
		log.Errorf(ctx, "Error when export: %v", err)
		panic(err)
	}
	log.Infof(ctx, "Exported file with %d records", total)
}
