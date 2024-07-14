package main

import (
	"context"
	"fmt"

	"go-service/internal/app"
)

func main() {
	var cfg app.Config
	cfg.Sql.Driver = "postgres"
	cfg.Sql.DataSourceName = "postgres://postgres:abcd1234@localhost/masterdata?sslmode=disable"

	ctx := context.Background()
	fmt.Println("Export file")
	app, err := app.NewApp(ctx, cfg)
	if err != nil {
		fmt.Println(ctx, "Error when initialize: "+err.Error())
		panic(err)
	}
	total, err := app.Export(ctx)
	if err != nil {
		fmt.Println("Error when export: " + err.Error())
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Exported file with %d records", total))
}
