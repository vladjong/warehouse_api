package csv

import (
	"fmt"
	"os"

	"github.com/adough/warehouse_api/internal/entity"
	"github.com/gocarina/gocsv"
)

type parser struct {
}

func New() *parser {
	return &parser{}
}

func (p *parser) GetWarehouses(filename string) ([]entity.Warehouse, error) {
	warehouseFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("[parser.GetWarehouse]:%v", err)
	}
	defer warehouseFile.Close()

	warehouses := []entity.Warehouse{}
	if err := gocsv.UnmarshalFile(warehouseFile, &warehouses); err != nil {
		return nil, fmt.Errorf("[parser.GetWarehouse]:%v", err)
	}
	return warehouses, nil
}

func (p *parser) GetProducts(filename string) ([]entity.Product, error) {
	productsFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("[parser.GetProducts]:%v", err)
	}
	defer productsFile.Close()

	products := []entity.Product{}
	if err := gocsv.UnmarshalFile(productsFile, &products); err != nil {
		return nil, fmt.Errorf("[parser.GetProducts]:%v", err)
	}
	return products, nil
}

func (p *parser) GetProductInWarehouse(filename string) ([]entity.ProductInWarehouse, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("[parser.GetProductInWarehouse]:%v", err)
	}
	defer file.Close()

	productInWarehouse := []entity.ProductInWarehouse{}
	if err := gocsv.UnmarshalFile(file, &productInWarehouse); err != nil {
		return nil, fmt.Errorf("[parser.GetProductInWarehouse]:%v", err)
	}
	return productInWarehouse, nil
}
