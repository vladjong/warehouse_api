package repository

import (
	"context"

	"github.com/adough/warehouse_api/internal/entity"
)

type Repository interface {
	AddWarehouse(ctx context.Context, data entity.Warehouse) error
	GetAllWarehouse(ctx context.Context) ([]entity.Warehouse, error)
	AddProduct(ctx context.Context, data entity.Product) error
	GetAllProduct(ctx context.Context) ([]entity.Product, error)
	AddProductInWarehouse(ctx context.Context, data entity.ProductInWarehouse) error
	ReservationProduct(ctx context.Context, data entity.ProductInWarehouse) error
}
