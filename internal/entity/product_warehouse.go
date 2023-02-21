package entity

import (
	"fmt"

	"github.com/google/uuid"
)

type ProductInWarehouse struct {
	Id          int64  `json:"-" db:"id" goqu:"skipinsert" csv:"id"`
	ProductId   string `json:"product_id" db:"product_id" csv:"product_id"`
	WarehouseId int64  `json:"warehouse_id" db:"warehouse_id" csv:"warehouse_id"`
	QTY         int    `json:"qty" db:"qty" csv:"qty"`
}

type AllProductInWarehouse struct {
	Product []Product
}

func NewProductInWarehouse(productId string, warehouseId int64, qty int) ProductInWarehouse {
	return ProductInWarehouse{
		ProductId:   productId,
		WarehouseId: warehouseId,
		QTY:         qty,
	}
}

func ValidationProductInWarehouse(data ProductInWarehouse) error {
	if _, err := uuid.Parse(data.ProductId); err != nil {
		return fmt.Errorf("[entity.ValidationProductInWarehouse]:%v", err)
	} else if data.WarehouseId <= 0 {
		return fmt.Errorf("[entity.ValidationProductInWarehouse]:incorrect warehouse_id=%v", data.WarehouseId)
	} else if data.QTY <= 0 {
		return fmt.Errorf("[entity.ValidationProductInWarehouse]:incorrect qty=%v", data.WarehouseId)
	}
	return nil
}
