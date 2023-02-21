package entity_test

import (
	"testing"

	"github.com/adough/warehouse_api/internal/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateProtuct(t *testing.T) {
	type testCase struct {
		name     string
		input    entity.Product
		hasError bool
	}

	testTable := []testCase{
		{
			name:     "correct struct",
			input:    entity.Product{uuid.NewString(), "test", 10, 10},
			hasError: false,
		},
		{
			name:     "incorrect name",
			input:    entity.Product{uuid.NewString(), "", 10, 10},
			hasError: true,
		},
		{
			name:     "incorrect size",
			input:    entity.Product{uuid.NewString(), "test", -10, 10},
			hasError: true,
		},
		{
			name:     "incorrect qty",
			input:    entity.Product{uuid.NewString(), "test", 10, -10},
			hasError: true,
		},
	}

	for _, test := range testTable {
		err := entity.ValidateProduct(test.input)
		if test.hasError {
			assert.NotNil(t, err, test.name)
		} else {
			assert.Nil(t, err, test.name)
		}
	}
}
