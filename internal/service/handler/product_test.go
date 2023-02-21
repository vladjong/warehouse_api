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
	mock_repository "github.com/adough/warehouse_api/internal/repository/mock"
	"github.com/adough/warehouse_api/internal/service/handler"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/stretchr/testify/assert"
)

type Response struct {
	Result *handler.Response `json:"result"`
	Error  interface{}       `json:"error"`
}

func init_rpc(handler handler.Service) *chi.Mux {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	rpcServer.RegisterService(handler, "warehouse")
	r := chi.NewRouter()
	r.Handle("/warehouse", rpcServer)

	return r
}

func TestAddProduct(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockRepository, data entity.Product)

	type testCase struct {
		name               string
		input              entity.Product
		inputBody          string
		mockBehavior       mockBehavior
		hasError           bool
		expectedResultBody string
	}

	testTable := []testCase{
		{
			name:  "correct",
			input: entity.Product{Id: "992afd25-09bd-49e6-82de-c873923d8d09", Name: "test_3", Size: 30, QTY: 30},
			inputBody: `{
				"jsonrpc": "2.0",
				"id": 2,
				"method": "warehouse.AddProduct",
				"params": [
					{
					"products": [
						{
							"Id":"992afd25-09bd-49e6-82de-c873923d8d09",
							"Name":"test_3",
							"Size": 30,
							"QTY": 30
						}
					]
				}
				]
			}`,
			mockBehavior: func(r *mock_repository.MockRepository, data entity.Product) {
				r.EXPECT().AddProduct(context.Background(), data).Return(nil).AnyTimes()
			},
			hasError:           false,
			expectedResultBody: "Successfully added",
		},
		{
			name:  "incorrect",
			input: entity.Product{Id: "992afd25-09bd-49e6-82de-c873923d8d09", Name: "test_3", Size: 30, QTY: 30},
			inputBody: `{
				"jsonrpc": "2.0",
				"id": 2,
				"method": "warehouse.AddProduct",
				"params": [
					{
					"products": [
						{
							"Id":"992afd25-09bd-49e6-82de-c873923d8d09",
							"Name":"test_3",
							"Size": 30,
							"QTY": 30
						}
					]
				}
				]
			}`,
			mockBehavior: func(r *mock_repository.MockRepository, data entity.Product) {
				r.EXPECT().AddProduct(context.Background(), data).Return(fmt.Errorf("error")).AnyTimes()
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
