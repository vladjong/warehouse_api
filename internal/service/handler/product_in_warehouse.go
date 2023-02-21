package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/adough/warehouse_api/internal/entity"
)

func (s *service) AddProductInWarehouse(r *http.Request, data *ProductInWarehouseArgs, response *Response) error {
	ctx := context.Background()
	if len(data.ProductInWarehouse) == 0 {
		log.Printf("[service.AddProductInWarehouse]:empty params")
		return fmt.Errorf("empty params")
	}
	for _, val := range data.ProductInWarehouse {
		if err := entity.ValidationProductInWarehouse(val); err != nil {
			log.Printf("[service.AddProductInWarehouse]:%v", err)
			return fmt.Errorf("[service.AddProductsInWarehouses]:%v", err)
		}
		if err := s.rep.AddProductInWarehouse(ctx, val); err != nil {
			log.Printf("[service.AddProductInWarehouse]:%v", err)
			return fmt.Errorf("[service.AddProductsInWarehouses]:%v", err)
		}
	}
	response.Status = "Successfully added"
	return nil
}

func (s *service) AddProductInWarehouseInFile(r *http.Request, data *FilenameArgs, response *Response) error {
	ctx := context.Background()
	if len(data.Filename) == 0 {
		log.Printf("[service.AddProductInWarehouseInFile]:empty params")
		return fmt.Errorf("empty params")
	}

	productInWarehouse, err := s.parser.GetProductInWarehouse(data.Filename)
	if err != nil {
		log.Printf("[service.AddProductInWarehouseInFile]:%v", err)
		return fmt.Errorf("[service.AddProductInWarehouseInFile]:%v", err)
	}

	for _, val := range productInWarehouse {
		if err := entity.ValidationProductInWarehouse(val); err != nil {
			log.Printf("[service.AddProductInWarehouseInFile]:%v", err)
			return fmt.Errorf("[service.AddProductInWarehouseInFile]:%v", err)
		}
		if err := s.rep.AddProductInWarehouse(ctx, val); err != nil {
			log.Printf("[service.AddProductInWarehouseInFile]:%v", err)
			return fmt.Errorf("[service.AddProductInWarehouseInFile]:%v", err)
		}
	}
	response.Status = "Successfully added"
	return nil
}

func (s *service) GetAllProductInWarehouse(r *http.Request, data *WarehouseIdArgs, response *ResponseProduct) error {
	ctx := context.Background()
	if data.WarehouseId <= 0 {
		log.Printf("[service.GetAllProductInWarehouse]:incorrect warehouse id")
		return fmt.Errorf("incorrect warehouse id")
	}
	products, err := s.rep.GetAllProductInWarehouse(ctx, data.WarehouseId)
	if err != nil {
		log.Printf("[service.GetAllProductInWarehouse]:%v", err)
		return fmt.Errorf("[service.GetAllProductInWarehouse]:%v", err)
	}
	response.Products = products
	response.Status = "Successfully"
	return nil
}
