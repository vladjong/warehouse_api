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

func TestAddWarehouse(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockRepository, data entity.Warehouse)

	type testCase struct {
		name               string
		input              entity.Warehouse
		inputBody          string
		mockBehavior       mockBehavior
		hasError           bool
		expectedResultBody string
	}

	testTable := []testCase{
		{
			name:  "correct",
			input: entity.Warehouse{Id: 1, Name: "test_1", IsAvalible: true},
			inputBody: `{
				"jsonrpc": "2.0",
				"id": 2,
				"method": "warehouse.AddWarehouses",
				"params": [
					{
					"warehouses": [
						{
							"id":"1",
							"name":"test_1_w",
							"is_avalible": true
						}
					]
				}
				]
			}`,
			mockBehavior: func(r *mock_repository.MockRepository, data entity.Warehouse) {
				r.EXPECT().AddWarehouse(context.Background(), data).Return(nil).AnyTimes()
			},
			hasError:           false,
			expectedResultBody: "Successfully added",
		},
		{
			name:  "incorrect",
			input: entity.Warehouse{Id: 1, Name: "test_1", IsAvalible: true},
			inputBody: `{
			"jsonrpc": "2.0",
			"id": 2,
			"method": "warehouse.AddWarehouses",
			"params": [
				{
				"warehouses": [
					{
						"id":"1",
						"name":"test_1_w",
						"is_avalible": true
					}
				]
			}
			]
		}`,
			mockBehavior: func(r *mock_repository.MockRepository, data entity.Warehouse) {
				r.EXPECT().AddWarehouse(context.Background(), data).Return(fmt.Errorf("error")).AnyTimes()
			},
			hasError:           true,
			expectedResultBody: "",
		},
	}

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
