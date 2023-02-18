package entity

type ProductReservation struct {
	Id                 int64 `db:"id" goqu:"skipinsert"`
	WarehouseProductId int64 `db:"id_warehouse_product"`
	QTY                int   `db:"qty"`
}
