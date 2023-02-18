package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adough/warehouse_api/internal/config"
	"github.com/adough/warehouse_api/internal/db"
	// "github.com/adough/warehouse_api/internal/entity"
	"github.com/adough/warehouse_api/internal/repository/postgres"
)

type Arith struct {
}

func main() {
	ctx := context.Background()
	cfg, err := config.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	pgx, err := db.NewPgx(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	postgres.New(pgx)
}
