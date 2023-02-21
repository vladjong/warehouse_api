package entity_test

import (
	"testing"

	"github.com/adough/warehouse_api/internal/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateProductInWarehouse(t *testing.T) {
	type testCase struct {
		name     string
		input    entity.ProductInWarehouse
		hasError bool
	}

	testTable := []testCase{
		{
			name:     "correct struct",
			input:    entity.ProductInWarehouse{1, uuid.NewString(), 1, 10},
			hasError: false,
		},
		{
			name:     "incorrect uuid",
			input:    entity.ProductInWarehouse{1, "123", 1, 10},
			hasError: true,
		},
		{
			name:     "incorrect product_id",
			input:    entity.ProductInWarehouse{1, uuid.NewString(), -1, 10},
			hasError: true,
		},
		{
			name:     "incorrect qty",
			input:    entity.ProductInWarehouse{1, uuid.NewString(), 1, -10},
			hasError: true,
		},
	}

	for _, test := range testTable {
		err := entity.ValidationProductInWarehouse(test.input)
		if test.hasError {
			assert.NotNil(t, err, test.name)
		} else {
			assert.Nil(t, err, test.name)
		}
	}
}
