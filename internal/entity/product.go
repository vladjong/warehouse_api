package entity

import (
	"fmt"
)

type Product struct {
	Id   string `json:"-" db:"id" csv:"id"`
	Name string `json:"name" db:"name" csv:"name"`
	Size int    `json:"size" db:"size" csv:"size"`
	QTY  int    `json:"qty" db:"qty" csv:"qty"`
}

func ValidateProduct(data Product) error {
	if len(data.Name) == 0 {
		return fmt.Errorf("[entiry.ValidateProduct]:len(Name)=0")
	} else if data.Size <= 0 {
		return fmt.Errorf("[entiry.ValidateProduct]:size=%v", data.Size)
	} else if data.QTY <= 0 {
		return fmt.Errorf("[entiry.ValidateProduct]:qty=%v", data.QTY)
	}
	return nil
}
