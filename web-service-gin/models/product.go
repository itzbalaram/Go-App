package models

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

// product represents data about a record product.
type Product struct {
	gorm.Model // adds ID, created_at etc.

	// ProductId    uint    `json:"productid" gorm:"primaryKey"`
	Name         string  `json:"name" gorm:"index"`
	Category     string  `json:"category"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
	LastModified string  `json:"lastModifed"`
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("invalid data for the field:Name")
	}
	if p.Category == "" {
		return fmt.Errorf("invalid data for the field:Category")
	}
	if p.Price == 0.00 {
		return fmt.Errorf("invalid data for the field:Price")
	}
	return nil
}

func (p *Product) ToJsonByte() ([]byte, error) {
	return json.Marshal(p)
}

// products slice to seed record product data.
var Products = []Product{
	{Name: "Redmi K20 Pro", Category: "Smartphone", Status: "Active", Price: 29999.99},
	{Name: "IPhone 13 pro", Category: "Smartphone", Status: "Active", Price: 75000.99},
	{Name: "HP-Pavilion 14", Category: "Laptop", Status: "Active", Price: 65000.99},
	{Name: "HP-Probook 16", Category: "Laptop", Status: "Active", Price: 85000.99},
}
