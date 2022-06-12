package main

import (
	"context"
	"fmt"
	"github.com/core-go/config"

	"go-service/internal/app"
)

func main() {
	var conf app.Config
	err := config.Load(&conf, "configs/config")
	if err != nil {
		panic(err)
	}

	fmt.Println("Export file")
	ctx := context.Background()
	app, err := app.NewApp(ctx, conf)
	if err != nil {
		panic(err)
	}
	err = app.Export(ctx)
	if err != nil {
		panic(err)
	}
}
