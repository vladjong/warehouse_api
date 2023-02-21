package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/adough/warehouse_api/internal/entity"
)

func (s *service) AddWarehouses(r *http.Request, data *WarehouseArgs, response *Response) error {
	ctx := context.Background()
	if len(data.Warehouse) == 0 {
		log.Printf("[service.AddWarehouse]:empty params")
		return fmt.Errorf("empty params")
	}
	for _, warhouse := range data.Warehouse {
		if err := entity.ValidateWarehouse(warhouse); err != nil {
			log.Printf("[service.AddWarehouse]:%v", err)
			return fmt.Errorf("[service.AddWarehouse]:%v", err)
		}
		if err := s.rep.AddWarehouse(ctx, warhouse); err != nil {
			log.Printf("[service.AddWarehouse]:%v", err)
			return fmt.Errorf("[service.AddWarehouse]:%v", err)
		}
	}
	response.Status = "Successfully added"
	return nil
}

func (s *service) AddWarehousesInFile(r *http.Request, data *FilenameArgs, response *Response) error {
	ctx := context.Background()
	if len(data.Filename) == 0 {
		log.Printf("[service.AddWarehousesInFile]:empty params")
		return fmt.Errorf("empty params")
	}
	warehouses, err := s.parser.GetWarehouses(data.Filename)
	if err != nil {
		log.Printf("[service.AddWarehousesInFile]:%v", err)
		return fmt.Errorf("[service.AddWarehousesInFile]:%v", err)
	}
	for _, warhouse := range warehouses {
		if err := entity.ValidateWarehouse(warhouse); err != nil {
			log.Printf("[service:AddWarehousesInFile]:%v", err)
			return fmt.Errorf("[service.AddWarehousesInFile]:%v", err)
		}
		if err := s.rep.AddWarehouse(ctx, warhouse); err != nil {
			log.Printf("[service.AddWarehousesInFile]:%v", err)
			return fmt.Errorf("[service.AddWarehousesInFile]:%v", err)
		}
	}
	response.Status = "Successfully added"
	return nil
}
