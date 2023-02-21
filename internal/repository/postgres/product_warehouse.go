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

	query, _, err := goqu.Insert(WAREHOUSE_PRODUCT).
		Rows(data).
		OnConflict(goqu.DoUpdate("ON CONSTRAINT unique_warehouse_product",
			goqu.Record{
				"qty": goqu.L(fmt.Sprintf("EXCLUDED.qty+%v.qty", WAREHOUSE_PRODUCT)),
			})).
		ToSQL()
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

func (r *repository) GetAllProductInWarehouse(ctx context.Context, id int64) ([]entity.Product, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("[postgres.GetAllProductInWarehouse]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	queryWarehouse, _, err := goqu.From(WAREHOUSE_TABLE).Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("[postgres.GetAllProductInWarehouse]:%v", err)
	}
	warehouse := entity.Warehouse{}
	if err := tx.GetContext(ctx, &warehouse, queryWarehouse); err != nil {
		return nil, fmt.Errorf("[postgres.GetAllProductInWarehouse]:%v", err)
	}
	if !warehouse.IsAvalible {
		return nil, fmt.Errorf("[postgres.GetAllProductInWarehouse]:warehouse=%v not avalible", warehouse.Id)
	}

	queryProduct, _, err := goqu.From(WAREHOUSE_PRODUCT).Join(
		goqu.T(PRODUCT_TABLE),
		goqu.On(goqu.Ex{
			fmt.Sprintf("%v.product_id", WAREHOUSE_PRODUCT): goqu.I(fmt.Sprintf("%v.id", PRODUCT_TABLE)),
		})).
		Select(
			fmt.Sprintf("%v.id", PRODUCT_TABLE),
			fmt.Sprintf("%v.name", PRODUCT_TABLE),
			fmt.Sprintf("%v.size", PRODUCT_TABLE),
			fmt.Sprintf("%v.qty", WAREHOUSE_PRODUCT),
		).
		Where(goqu.C("warehouse_id").Eq(id)).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("[postgres.GetAllProductInWarehouse]:%v", err)
	}
	fmt.Println(queryProduct)
	products := []entity.Product{}
	if err := tx.SelectContext(ctx, &products, queryProduct); err != nil {
		return nil, fmt.Errorf("[postgres.GetAllProductInWarehouse]:%v", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("[postgres.GetAllProductInWarehouse]:%v", err)
	}
	return products, nil

}
