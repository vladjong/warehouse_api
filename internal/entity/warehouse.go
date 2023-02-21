package entity

import "fmt"

type Warehouse struct {
	Id         int64  `json:"id" db:"id" goqu:"skipinsert" csv:"id"`
	Name       string `json:"name" db:"name" csv:"name"`
	IsAvalible bool   `json:"is_avalible" db:"is_available" csv:"is_avalible"`
}

func ValidateWarehouse(data Warehouse) error {
	if len(data.Name) == 0 {
		return fmt.Errorf("[entiry.ValidateWarehouse]:len(Name)=0")
	}
	return nil
}
