package parser

import "github.com/adough/warehouse_api/internal/entity"

type Parser interface {
	GetWarehouses(filename string) ([]entity.Warehouse, error)
	GetProducts(filename string) ([]entity.Product, error)
	GetProductInWarehouse(filename string) ([]entity.ProductInWarehouse, error)
}
