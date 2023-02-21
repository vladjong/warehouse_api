package handler

import "github.com/adough/warehouse_api/internal/entity"

type ProductArgs struct {
	Products []entity.Product `json:"products"`
}

type WarehouseArgs struct {
	Warehouse []entity.Warehouse `json:"warehouses"`
}

type ProductWarehouse struct {
	ProductId   string `json:"product_id"`
	WarehouseId int64  `json:"warehouse_id"`
}

type ProductWarehouseArgs struct {
	ProductInWarehouse []ProductWarehouse `json:"products_warehouses"`
}

type ProductInWarehouseArgs struct {
	ProductInWarehouse []entity.ProductInWarehouse `json:"products_in_warehouses"`
}

type WarehouseIdArgs struct {
	WarehouseId int64 `json:"warehouse_id"`
}

type FilenameArgs struct {
	Filename string `json:"filename"`
}

type Response struct {
	Status string `json:"status"`
}

type ResponseProduct struct {
	Status   string           `json:"status"`
	Products []entity.Product `json:"products"`
}
