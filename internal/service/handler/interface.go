package handler

import "net/http"

type Service interface {
	AddProduct(r *http.Request, data *ProductArgs, response *Response) error
	AddProductInWarehouse(r *http.Request, data *ProductInWarehouseArgs, response *Response) error
	AddWarehouses(r *http.Request, data *WarehouseArgs, response *Response) error
	ReservatuinProduct(r *http.Request, data *ProductWarehouseArgs, response *Response) error
	GetAllProductInWarehouse(r *http.Request, data *WarehouseIdArgs, response *ResponseProduct) error
	RealeaseOfReserve(r *http.Request, data *ProductWarehouseArgs, response *Response) error
	AddProductInWarehouseInFile(r *http.Request, data *FilenameArgs, response *Response) error
	AddWarehousesInFile(r *http.Request, data *FilenameArgs, response *Response) error
}
