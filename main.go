package main

import (
	"context"
	"github.com/core-go/config"
	"github.com/core-go/log/rotatelogs"
	log "github.com/core-go/log/zap"

	"go-service/internal/app"
)

func main() {
	var conf app.Config
	err := config.Load(&conf, "configs/config")
	if err != nil {
		panic(err)
	}
	_, err = log.InitializeWithWriter(conf.Log, rotatelogs.GetWriter)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	log.Info(ctx, "Export file")
	app, err := app.NewApp(ctx, conf)
	if err != nil {
		log.Errorf(ctx, "Error when initialize: %v", err)
		panic(err)
	}
	err = app.Export(ctx)
	if err != nil {
		log.Errorf(ctx, "Error when export: %v", err)
		panic(err)
	}
	log.Info(ctx, "Exported file")
}
