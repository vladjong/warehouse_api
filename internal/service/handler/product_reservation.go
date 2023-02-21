package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/adough/warehouse_api/internal/entity"
)

func (s *service) ReservatuinProduct(r *http.Request, data *ProductWarehouseArgs, response *Response) error {
	ctx := context.Background()
	if len(data.ProductInWarehouse) == 0 {
		log.Printf("[service.ReservationProduct:empty params")
		return fmt.Errorf("empty params")
	}
	mapProduct := make(map[ProductWarehouse]int)
	for _, val := range data.ProductInWarehouse {
		mapProduct[val] += 1
	}
	for key, val := range mapProduct {
		productInWarehouse := entity.NewProductInWarehouse(key.ProductId, key.WarehouseId, val)
		if err := entity.ValidationProductInWarehouse(productInWarehouse); err != nil {
			log.Printf("[service.ReservationProduct]:%v", err)
			return fmt.Errorf("[service.ReservationProduct]:%v", err)
		}
		if err := s.rep.ReservationProduct(ctx, productInWarehouse); err != nil {
			log.Printf("[service.ReservationProduct]:%v", err)
			return fmt.Errorf("[service.ReservationProduct]:%v", err)
		}
	}
	response.Status = "Successfully reservation"
	return nil
}

func (s *service) RealeaseOfReserve(r *http.Request, data *ProductWarehouseArgs, response *Response) error {
	ctx := context.Background()
	if len(data.ProductInWarehouse) == 0 {
		log.Printf("[service.RealeaseOfReserve]:empty params")
		return fmt.Errorf("empty params")
	}
	mapProduct := make(map[ProductWarehouse]int)
	for _, val := range data.ProductInWarehouse {
		mapProduct[val] += 1
	}
	for key, val := range mapProduct {
		productInWarehouse := entity.NewProductInWarehouse(key.ProductId, key.WarehouseId, val)
		if err := entity.ValidationProductInWarehouse(productInWarehouse); err != nil {
			log.Printf("[service.RealeaseOfReserve]:%v", err)
			return fmt.Errorf("[service.RealeaseOfReserve]:%v", err)
		}
		if err := s.rep.RealeaseOfReserve(ctx, productInWarehouse); err != nil {
			log.Printf("[service.RealeaseOfReserve]:%v", err)
			return fmt.Errorf("[service.RealeaseOfReserve]:%v", err)
		}
	}
	response.Status = "Successfully realease"
	return nil
}
