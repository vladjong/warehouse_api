package main

import (
	"context"
	"log"

	"github.com/adough/warehouse_api/internal/app"
	"github.com/adough/warehouse_api/internal/config"
	"github.com/adough/warehouse_api/internal/db"
	"github.com/adough/warehouse_api/internal/parser/csv"
	"github.com/adough/warehouse_api/internal/repository/postgres"
	"github.com/adough/warehouse_api/internal/service/handler"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("init config")

	pgx, err := db.NewPgx(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("init pgx driver")

	repository := postgres.New(pgx)

	if err := db.Migrate(cfg.DB); err != nil {
		log.Fatal(err)
	}
	log.Println("completed migrate")

	parser := csv.New()

	handlerWarehouse := handler.NewServie(repository, parser)

	service := app.New(handlerWarehouse, *cfg)
	service.Start()
}
