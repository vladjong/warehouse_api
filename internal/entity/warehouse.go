package entity

type Warehouse struct {
	Id         int64  `db:"id" goqu:"skipinsert"`
	Name       string `db:"name"`
	IsAvalible bool   `db:"is_available"`
}
