package handler_test

import (
	"bytes"
	"context"
	json2 "encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/adough/warehouse_api/internal/entity"
	mock_parser "github.com/adough/warehouse_api/internal/parser/mock"
	"github.com/adough/warehouse_api/internal/service/handler"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_repository "github.com/adough/warehouse_api/internal/repository/mock"
)

type ResponseProduct struct {
	Result *handler.ResponseProduct `json:"result"`
	Error  interface{}              `json:"error"`
}

func TestAddProductInWarehouse(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockRepository, data entity.ProductInWarehouse)

	type testCase struct {
		name               string
		input              entity.ProductInWarehouse
		inputBody          string
		mockBehavior       mockBehavior
		hasError           bool
		expectedResultBody string
	}

	testTable := []testCase{
		{
			name:  "correct",
			input: entity.ProductInWarehouse{Id: 1, ProductId: "992afd25-09bd-49e6-82de-c873923d8d09", WarehouseId: 1, QTY: 10},
			inputBody: `{
				"jsonrpc": "2.0",
				"id": 3,
				"method": "warehouse.AddProductInWarehouse",
				"params": [
					{
					"products_in_warehouses": [
						{
							"id":1.
							"product_id": "992afd25-09bd-49e6-82de-c873923d8d09",
							"warehouse_id": 1,
							"qty": 10
						}
					]
				}
				]
			}`,
			mockBehavior: func(r *mock_repository.MockRepository, data entity.ProductInWarehouse) {
				r.EXPECT().AddProductInWarehouse(context.Background(), data).Return(nil).AnyTimes()
			},
			hasError:           false,
			expectedResultBody: "Successfully added",
		},
		{
			name:  "incorrect",
			input: entity.ProductInWarehouse{Id: 1, ProductId: "992afd25-09bd-49e6-82de-c873923d8d09", WarehouseId: 1, QTY: 10},
			inputBody: `{
				"jsonrpc": "2.0",
				"id": 3,
				"method": "warehouse.AddProductInWarehouse",
				"params": [
					{
					"products_in_warehouses": [
						{
							"id":1.
							"product_id": "992afd25-09bd-49e6-82de-c873923d8d09",
							"warehouse_id": 1,
							"qty": 10
						}
					]
				}
				]
			}`,
			mockBehavior: func(r *mock_repository.MockRepository, data entity.ProductInWarehouse) {
				r.EXPECT().AddProductInWarehouse(context.Background(), data).Return(fmt.Errorf("error")).AnyTimes()
			},
			hasError:           false,
			expectedResultBody: "",
		}}

	for _, test := range testTable {
		c := gomock.NewController(t)
		defer c.Finish()
		rep := mock_repository.NewMockRepository(c)
		parser := mock_parser.NewMockParser(c)
		service := handler.NewServie(rep, parser)
		test.mockBehavior(rep, test.input)

		r := init_rpc(service)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/warehouse", bytes.NewBufferString(test.inputBody))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		response := Response{}
		if err := json2.Unmarshal(w.Body.Bytes(), &response); err != nil {
			return
		}
		if test.hasError {
			assert.NotNil(t, response.Error)
		} else {
			assert.Equal(t, test.expectedResultBody, response.Result.Status)
			assert.Nil(t, response.Error)
		}
	}
}

func TestGetAllProductInWarehouse(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockRepository, data int64)

	type testCase struct {
		name               string
		input              int64
		inputBody          string
		mockBehavior       mockBehavior
		hasError           bool
		expectedResultBody string
	}

	testTable := []testCase{
		{
			name:  "correct",
			input: 1,
			inputBody: `{
				"jsonrpc": "2.0",
				"id": 1,
				"method": "warehouse.GetAllProductInWarehouse",
				"params": [
					{
					"warehouse_id": 1
					}
				]
			}`,
			mockBehavior: func(r *mock_repository.MockRepository, data int64) {
				r.EXPECT().GetAllProductInWarehouse(context.Background(), data).Return([]entity.Product{{Id: "992afd25-09bd-49e6-82de-c873923d8d09", Name: "test_3", Size: 30, QTY: 30}}, nil).AnyTimes()
			},
			hasError:           false,
			expectedResultBody: "Successfully",
		},
		{
			name:  "incorrect",
			input: 1,
			inputBody: `{
				"jsonrpc": "2.0",
				"id": 1,
				"method": "warehouse.GetAllProductInWarehouse",
				"params": [
					{
					"warehouse_id": 1
					}
				]
			}`,
			mockBehavior: func(r *mock_repository.MockRepository, data int64) {
				r.EXPECT().GetAllProductInWarehouse(context.Background(), data).Return(nil, fmt.Errorf("error")).AnyTimes()
			},
			hasError:           true,
			expectedResultBody: "",
		}}

	for _, test := range testTable {
		c := gomock.NewController(t)
		defer c.Finish()
		rep := mock_repository.NewMockRepository(c)
		parser := mock_parser.NewMockParser(c)
		service := handler.NewServie(rep, parser)
		test.mockBehavior(rep, test.input)

		r := init_rpc(service)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/warehouse", bytes.NewBufferString(test.inputBody))
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		response := ResponseProduct{}
		if err := json2.Unmarshal(w.Body.Bytes(), &response); err != nil {
			return
		}
		if test.hasError {
			assert.NotNil(t, response.Error)
		} else {
			assert.Equal(t, test.expectedResultBody, response.Result.Status)
			assert.Nil(t, response.Error)
		}
	}
}
