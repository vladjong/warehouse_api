package entity_test

import (
	"testing"

	"github.com/adough/warehouse_api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestValidateWarehouse(t *testing.T) {
	type testCase struct {
		name     string
		input    entity.Warehouse
		hasError bool
	}

	testTable := []testCase{
		{
			name:     "correct struct",
			input:    entity.Warehouse{1, "test", true},
			hasError: false,
		},
		{
			name:     "incorrect name",
			input:    entity.Warehouse{1, "", true},
			hasError: true,
		},
	}

	for _, test := range testTable {
		err := entity.ValidateWarehouse(test.input)
		if test.hasError {
			assert.NotNil(t, err, test.name)
		} else {
			assert.Nil(t, err, test.name)
		}
	}
}
