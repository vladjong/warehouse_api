package repository

import (
	"context"

	"github.com/adough/warehouse_api/internal/entity"
)

//go:generate mockgen -source=interface.go -destination=mock/mock.go

type Repository interface {
	AddWarehouse(ctx context.Context, data entity.Warehouse) error
	GetAllWarehouse(ctx context.Context) ([]entity.Warehouse, error)
	AddProduct(ctx context.Context, data entity.Product) error
	GetAllProduct(ctx context.Context) ([]entity.Product, error)
	AddProductInWarehouse(ctx context.Context, data entity.ProductInWarehouse) error
	ReservationProduct(ctx context.Context, data entity.ProductInWarehouse) error
	GetAllProductInWarehouse(ctx context.Context, id int64) ([]entity.Product, error)
	RealeaseOfReserve(ctx context.Context, data entity.ProductInWarehouse) error
}
