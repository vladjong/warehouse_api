package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/adough/warehouse_api/internal/entity"
	"github.com/google/uuid"
)

func (s *service) AddProduct(r *http.Request, data *ProductArgs, response *Response) error {
	ctx := context.Background()
	if len(data.Products) == 0 {
		log.Printf("[service.AddProducts]:empty params")
		return fmt.Errorf("empty params")
	}
	for _, product := range data.Products {
		if err := entity.ValidateProduct(product); err != nil {
			log.Printf("[service.AddProduct]:%v", err)
			return fmt.Errorf("[service.AddProduct]:%v", err)
		}
		product.Id = uuid.NewString()
		if err := s.rep.AddProduct(ctx, product); err != nil {
			log.Printf("[service.AddProduct]:%v", err)
			return fmt.Errorf("[service.AddProduct]:%v", err)
		}
	}
	response.Status = "Successfully added"
	return nil
}

func (s *service) AddProductsInFile(r *http.Request, data *FilenameArgs, response *Response) error {
	ctx := context.Background()
	if len(data.Filename) == 0 {
		log.Printf("[service.AddProductsInFile]:empty params")
		return fmt.Errorf("empty params")
	}

	products, err := s.parser.GetProducts(data.Filename)
	if err != nil {
		log.Printf("[service.AddProductsInFile]:%v", err)
		return fmt.Errorf("[service.AddProductsInFile]:%v", err)
	}

	for _, product := range products {
		if err := entity.ValidateProduct(product); err != nil {
			log.Printf("[service.AddProductsInFile]:%v", err)
			return fmt.Errorf("[service.AddProductsInFile]:%v", err)
		}
		if err := s.rep.AddProduct(ctx, product); err != nil {
			log.Printf("[service.AddProductsInFile]:%v", err)
			return fmt.Errorf("[service.AddProductsInFile]:%v", err)
		}
	}
	response.Status = "Successfully added"
	return nil
}
