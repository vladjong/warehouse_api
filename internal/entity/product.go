package entity

type Product struct {
	Id   string `db:"id"`
	Name string `db:"name"`
	Size int    `db:"size"`
	QTY  int    `db:"qty"`
}
