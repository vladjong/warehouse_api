package postgres

import (
	"context"
	"fmt"

	"github.com/adough/warehouse_api/internal/entity"
	"github.com/doug-martin/goqu/v9"
)

const (
	PRODUCT_RESERVATION_TABLE = "product_reservation"
)

func (r *repository) ReservationProduct(ctx context.Context, data entity.ProductInWarehouse) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("[postgres.ReservationProduct]:%v", err)
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

	product := entity.ProductInWarehouse{}
	queryProductWarehouse, _, err := goqu.From(WAREHOUSE_PRODUCT).
		Where(goqu.C("product_id").Eq(data.ProductId)).
		Where(goqu.C("warehouse_id").Eq(data.WarehouseId)).
		ToSQL()
	if err != nil {
		return fmt.Errorf("[postgres.ReservationProduct]:%v", err)
	}
	if err := tx.GetContext(ctx, &product, queryProductWarehouse); err != nil {
		return fmt.Errorf("[postgres.ReservationProduct]:%v", err)
	}
	if product.QTY < data.QTY {
		return fmt.Errorf("[postgres.ReservationProduct]:product=%v in warehouse less then need qty=%v", product.QTY, data.QTY)
	}

	queryUpdateProductWarehouse, _, err := goqu.Update(WAREHOUSE_PRODUCT).
		Set(goqu.Record{"qty": goqu.L(fmt.Sprintf("%v.qty-%v", WAREHOUSE_PRODUCT, data.QTY))}).
		Where(goqu.C("product_id").Eq(data.ProductId)).
		Where(goqu.C("warehouse_id").Eq(data.WarehouseId)).
		ToSQL()
	fmt.Println(queryUpdateProductWarehouse)
	if err != nil {
		return fmt.Errorf("[postgres.ReservationProduct]:%v", err)
	}
	if _, err := tx.ExecContext(ctx, queryUpdateProductWarehouse); err != nil {
		return fmt.Errorf("[postgres.ReservationProduct]:%v", err)
	}

	queryUpdateProduct, _, err := goqu.Update(PRODUCT_TABLE).
		Set(goqu.Record{"qty": goqu.L(fmt.Sprintf("%v.qty-%v", PRODUCT_TABLE, data.QTY))}).
		Where(goqu.C("id").Eq(data.ProductId)).
		ToSQL()
	if err != nil {
		return fmt.Errorf("[postgres.ReservationProduct]:%v", err)
	}
	if _, err := tx.ExecContext(ctx, queryUpdateProduct); err != nil {
		return fmt.Errorf("[postgres.ReservationProduct]:%v", err)
	}

	insert := entity.ProductReservation{
		WarehouseProductId: product.Id,
		QTY:                data.QTY,
	}
	queryInsert, _, err := goqu.Insert(PRODUCT_RESERVATION_TABLE).
		Rows(insert).
		OnConflict(goqu.DoUpdate("id_warehouse_product",
			goqu.Record{
				"qty": goqu.L(fmt.Sprintf("EXCLUDED.qty+%v.qty", PRODUCT_RESERVATION_TABLE)),
			})).
		ToSQL()
	if err != nil {
		return fmt.Errorf("[postgres.ReservationProduct]:%v", err)
	}
	if _, err := tx.ExecContext(ctx, queryInsert); err != nil {
		return fmt.Errorf("[postgres.ReservationProduct]:%v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgres.ReservationProduct]:%v", err)
	}
	return nil
}

func (r *repository) RealeaseOfReserve(ctx context.Context, data entity.ProductInWarehouse) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("[postgres.RealeaseOfReserve]:%v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	reserve := entity.ProductReservation{}
	queryReserve, _, err := goqu.From(PRODUCT_RESERVATION_TABLE).Join(
		goqu.T(WAREHOUSE_PRODUCT),
		goqu.On(goqu.Ex{
			fmt.Sprintf("%v.id_warehouse_product", PRODUCT_RESERVATION_TABLE): goqu.I(fmt.Sprintf("%v.id", WAREHOUSE_PRODUCT)),
		})).
		Select(
			fmt.Sprintf("%v.id", PRODUCT_RESERVATION_TABLE),
			fmt.Sprintf("%v.id_warehouse_product", PRODUCT_RESERVATION_TABLE),
			fmt.Sprintf("%v.qty", PRODUCT_RESERVATION_TABLE),
		).
		Where(goqu.C("product_id").Eq(data.ProductId)).
		Where(goqu.C("warehouse_id").Eq(data.WarehouseId)).
		ToSQL()
	if err != nil {
		return fmt.Errorf("[postgres.RealeaseOfReserve]:%v", err)
	}
	if err := tx.GetContext(ctx, &reserve, queryReserve); err != nil {
		return fmt.Errorf("[postgres.RealeaseOfReserve]:%v", err)
	}
	if reserve.QTY < data.QTY {
		return fmt.Errorf("[postgres.RealeaseOfReserve]:reserve=%v in warehouse less then need qty=%v", reserve.QTY, data.QTY)
	}

	queryUpdateReserve, _, err := goqu.Update(PRODUCT_RESERVATION_TABLE).
		Set(goqu.Record{"qty": goqu.L(fmt.Sprintf("%v.qty-%v", PRODUCT_RESERVATION_TABLE, data.QTY))}).
		Where(goqu.C("id").Eq(reserve.Id)).
		ToSQL()
	if err != nil {
		return fmt.Errorf("[postgres.RealeaseOfReserve]:%v", err)
	}
	if _, err := tx.ExecContext(ctx, queryUpdateReserve); err != nil {
		return fmt.Errorf("[postgres.RealeaseOfReserve]:%v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgres.RealeaseOfReserve]:%v", err)
	}

	return nil
}
