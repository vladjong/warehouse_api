package main

import (
	"context"
	"log"

	"github.com/adough/warehouse_api/internal/config"
	"github.com/adough/warehouse_api/internal/controller/api"
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

	db := postgres.New(pgx)

	parser := csv.New()

	handlerWarehouse := handler.NewServie(db, parser)

	service := api.New(db, handlerWarehouse, *cfg)
	service.Start()
}
