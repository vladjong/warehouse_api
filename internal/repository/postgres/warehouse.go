package postgres

import (
	"context"
	"fmt"

	"github.com/adough/warehouse_api/internal/entity"
	"github.com/doug-martin/goqu/v9"
)

const (
	WAREHOUSE_TABLE = "warehouse"
)

func (r *repository) AddWarehouse(ctx context.Context, data entity.Warehouse) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("[postgres.AddWarehouse]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	query, _, err := goqu.Insert(WAREHOUSE_TABLE).Rows(data).OnConflict(goqu.DoNothing()).ToSQL()
	if err != nil {
		return fmt.Errorf("[postgres.AddWarehouse]:%v", err)
	}

	if _, err := tx.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("[postgres.AddWarehouse]:%v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgres.AddWarehouse]:%v", err)
	}
	return nil
}

func (r *repository) GetAllWarehouse(ctx context.Context) ([]entity.Warehouse, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("[postgres.GetAllWarehouse]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	query, _, err := goqu.From(WAREHOUSE_TABLE).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("[postgres.GetAllWarehouse]:%v", err)
	}

	data := []entity.Warehouse{}

	if err := tx.SelectContext(ctx, &data, query); err != nil {
		return nil, fmt.Errorf("[postgres.GetAllWarehouse]:%v", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("[postgres.GetAllWarehouse]:%v", err)
	}
	return data, nil
}
