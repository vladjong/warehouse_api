package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adough/warehouse_api/internal/config"
	"github.com/adough/warehouse_api/internal/db"
	"github.com/adough/warehouse_api/internal/entity"
	"github.com/adough/warehouse_api/internal/repository/postgres"
	"github.com/google/uuid"
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

	rep := postgres.New(pgx)

	data1 := entity.Warehouse{
		Name:       "MSK",
		IsAvalible: true,
	}

	data2 := entity.Warehouse{
		Name:       "Kem1",
		IsAvalible: true,
	}

	if err := rep.AddWarehouse(ctx, data1); err != nil {
		log.Fatal(err)
	}

	if err := rep.AddWarehouse(ctx, data2); err != nil {
		log.Fatal(err)
	}

	data, err := rep.GetAllWarehouse(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	product := entity.Product{
		Id:   uuid.NewString(),
		Name: "Test",
		Size: 20,
		QTY:  1,
	}

	if err := rep.AddProduct(ctx, product); err != nil {
		log.Fatal(err)
	}
	res, err := rep.GetAllProduct(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	da := entity.ProductInWarehouse{
		ProductId:   res[0].Id,
		WarehouseId: data[2].Id,
		QTY:         1,
	}

	if err := rep.AddProductInWarehouse(ctx, da); err != nil {
		log.Fatal(err)
	}

}
