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

func TestReservatuinProduct(t *testing.T) {
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
			input: entity.ProductInWarehouse{Id: 1, ProductId: "dd472911-a81c-49bb-9b8b-06d6186730f5", WarehouseId: 1, QTY: 10},
			inputBody: `{
				"jsonrpc": "2.0",
				"id": 1,
				"method": "warehouse.ReservatuinProduct",
				"params": [
					{
					"products_warehouses": [
						{
							"id":1
							"product_id":"dd472911-a81c-49bb-9b8b-06d6186730f5",
							"warehouse_id":1
						},
					]
				}
				]
			}`,
			mockBehavior: func(r *mock_repository.MockRepository, data entity.ProductInWarehouse) {
				r.EXPECT().ReservationProduct(context.Background(), data).Return(nil).AnyTimes()
			},
			hasError:           false,
			expectedResultBody: "Successfully added",
		},
		{
			name:  "incorrect",
			input: entity.ProductInWarehouse{Id: 1, ProductId: "992afd25-09bd-49e6-82de-c873923d8d09", WarehouseId: 1, QTY: 10},
			inputBody: `{
				"jsonrpc": "2.0",
				"id": 1,
				"method": "warehouse.ReservatuinProduct",
				"params": [
					{
					"products_warehouses": [
						{
							"id":1
							"product_id":"dd472911-a81c-49bb-9b8b-06d6186730f5",
							"warehouse_id":1
						},
					]
				}
				]
			}`,
			mockBehavior: func(r *mock_repository.MockRepository, data entity.ProductInWarehouse) {
				r.EXPECT().ReservationProduct(context.Background(), data).Return(fmt.Errorf("error")).AnyTimes()
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

func TestRealeaseProduct(t *testing.T) {
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
			input: entity.ProductInWarehouse{Id: 1, ProductId: "dd472911-a81c-49bb-9b8b-06d6186730f5", WarehouseId: 1, QTY: 10},
			inputBody: `{
				"jsonrpc": "2.0",
				"id": 1,
				"method": "warehouse.ReservatuinProduct",
				"params": [
					{
					"products_warehouses": [
						{
							"id":1
							"product_id":"dd472911-a81c-49bb-9b8b-06d6186730f5",
							"warehouse_id":1
						},
					]
				}
				]
			}`,
			mockBehavior: func(r *mock_repository.MockRepository, data entity.ProductInWarehouse) {
				r.EXPECT().RealeaseOfReserve(context.Background(), data).Return(nil).AnyTimes()
			},
			hasError:           false,
			expectedResultBody: "Successfully added",
		},
		{
			name:  "incorrect",
			input: entity.ProductInWarehouse{Id: 1, ProductId: "992afd25-09bd-49e6-82de-c873923d8d09", WarehouseId: 1, QTY: 10},
			inputBody: `{
				"jsonrpc": "2.0",
				"id": 1,
				"method": "warehouse.ReservatuinProduct",
				"params": [
					{
					"products_warehouses": [
						{
							"id":1
							"product_id":"dd472911-a81c-49bb-9b8b-06d6186730f5",
							"warehouse_id":1
						},
					]
				}
				]
			}`,
			mockBehavior: func(r *mock_repository.MockRepository, data entity.ProductInWarehouse) {
				r.EXPECT().RealeaseOfReserve(context.Background(), data).Return(fmt.Errorf("error")).AnyTimes()
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
