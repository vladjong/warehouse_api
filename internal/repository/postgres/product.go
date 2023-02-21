package postgres

import (
	"context"
	"fmt"

	"github.com/adough/warehouse_api/internal/entity"
	"github.com/doug-martin/goqu/v9"
)

const (
	PRODUCT_TABLE = "product"
)

func (r *repository) AddProduct(ctx context.Context, data entity.Product) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("[postgres.AddProduct]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	query, _, err := goqu.Insert(PRODUCT_TABLE).Rows(data).OnConflict(goqu.DoNothing()).ToSQL()
	if err != nil {
		return fmt.Errorf("[postgres.AddProduct]:%v", err)
	}

	if _, err := tx.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("[postgres.AddProduct]:%v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgres.AddWarehouse]:%v", err)
	}
	return nil
}

func (r *repository) GetAllProduct(ctx context.Context) ([]entity.Product, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("[postgres.GetAllProduct]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	query, _, err := goqu.From(PRODUCT_TABLE).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("[postgres.GetAllProduct]:%v", err)
	}

	data := []entity.Product{}

	if err := tx.SelectContext(ctx, &data, query); err != nil {
		return nil, fmt.Errorf("[postgres.GetAllProduct]:%v", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("[postgres.GetAllProduct]:%v", err)
	}
	return data, nil
}
