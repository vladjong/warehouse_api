package entity

type ProductInWarehouse struct {
	Id          int64  `db:"id" goqu:"skipinsert"`
	ProductId   string `db:"product_id"`
	WarehouseId int64  `db:"warehouse_id"`
	QTY         int    `db:"qty"`
}

type AllProductInWarehouse struct {
	Product []Product
}
