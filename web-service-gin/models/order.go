package models

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

// product represents data about a record product.
type Order struct {
	gorm.Model         // adds ID, created_at etc.
	CustomerId string  `json:"custid"`
	ProductId  string  `json:"prodid"`
	DeliveryBy string  `json:"deliveryby"`
	Amount     float64 `json:"amount"`
	// LastModified string `json:"lastModifed"`
}

func (c *Order) Validate() error {
	if c.ProductId == "" {
		return fmt.Errorf("invalid data for the field: ProductId")
	}
	if c.DeliveryBy == "" {
		return fmt.Errorf("invalid data for the field: deliveryby")
	}
	return nil
}

func (c *Order) ToJsonByte() ([]byte, error) {
	return json.Marshal(c)
}
