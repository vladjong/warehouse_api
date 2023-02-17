package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/adough/warehouse_api/internal/entity"
	"github.com/doug-martin/goqu/v9"
)

const (
	WAREHOUSE_PRODUCT = "warehouse_product"
)

func (r *repository) AddProductInWarehouse(ctx context.Context, data entity.ProductInWarehouse) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("[postgres.AddProductInWarehouse]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	queryProductSum, _, err := goqu.From(WAREHOUSE_PRODUCT).
		Select(goqu.SUM("qty")).
		GroupBy("product_id").
		Where(goqu.C("product_id").Eq(data.ProductId)).ToSQL()
	productQTY := 0
	if err := tx.GetContext(ctx, &productQTY, queryProductSum); err != nil {
		if err != sql.ErrNoRows {
			return fmt.Errorf("[postgres.AddProductInWarehouse]:%v", err)
		}
	}

	queryProduct, _, err := goqu.From(PRODUCT_TABLE).Where(goqu.C("id").Eq(data.ProductId)).ToSQL()
	if err != nil {
		return fmt.Errorf("[postgres.AddProductInWarehouse]:%v", err)
	}
	product := entity.Product{}
	if err := tx.GetContext(ctx, &product, queryProduct); err != nil {
		return fmt.Errorf("[postgres.AddProductInWarehouse]:%v", err)
	}
	if product.QTY < data.QTY+productQTY {
		return fmt.Errorf("[postgres.AddProductInWarehouse]:product.QTY=%v less then data.QTY=%v", product.QTY, data.QTY+productQTY)
	}

	queryWarehouse, _, err := goqu.From(WAREHOUSE_TABLE).Where(goqu.C("id").Eq(data.WarehouseId)).ToSQL()
	if err != nil {
		return fmt.Errorf("[postgres.AddProductInWarehouse]:%v", err)
	}
	warehouse := entity.Warehouse{}
	if err := tx.GetContext(ctx, &warehouse, queryWarehouse); err != nil {
		return fmt.Errorf("[postgres.AddProductInWarehouse]:%v", err)
	}
	if !warehouse.IsAvalible {
		return fmt.Errorf("[postgres.AddProductInWarehouse]:warehouse=%v not avalible", warehouse.Id)
	}

	query, _, err := goqu.Insert(WAREHOUSE_PRODUCT).
		Rows(data).
		OnConflict(goqu.DoUpdate("ON CONSTRAINT unique_warehouse_product",
			goqu.Record{
				"qty": goqu.L(fmt.Sprint("EXCLUDED.qty+warehouse_product.qty")),
			}),
		).ToSQL()
	if err != nil {
		return fmt.Errorf("[postgres.AddProductInWarehouse]:%v", err)
	}
	if _, err := tx.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("[postgres.AddProductInWarehouse]:%v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgres.AddProductInWarehouse]:%v", err)
	}
	return nil
}
