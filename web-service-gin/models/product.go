package models

import (
	"encoding/json"
	"fmt"
)

// product represents data about a record product.
type Product struct {
	ProductId    uint    `json:"productid" gorm:"primaryKey"`
	Name         string  `json:"name" gorm:"index"`
	Category     string  `json:"category"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
	LastModified string  `json:"lastModifed"`
}

func (c *Product) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("invalid data for the field:Name")
	}
	if c.Category == "" {
		return fmt.Errorf("invalid data for the field:Category")
	}
	if c.Price == 0.00 {
		return fmt.Errorf("invalid data for the field:Price")
	}
	return nil
}

func (c *Product) ToJsonByte() ([]byte, error) {
	return json.Marshal(c)
}

// products slice to seed record product data.
var Products = []Product{
	{Name: "Redmi K20 Pro", Category: "Smartphone", Status: "Active", Price: 29999.99},
	{Name: "IPhone 13 pro", Category: "Smartphone", Status: "Active", Price: 75000.99},
	{Name: "HP-Pavilion 14", Category: "Laptop", Status: "Active", Price: 65000.99},
	{Name: "HP-Probook 16", Category: "Laptop", Status: "Active", Price: 85000.99},
}

// type Contact struct {
// 	ID           uint   `json:"id" gorm:"primaryKey"`
// 	Name         string `json:"name" gorm:"index"`
// 	Address      string `json:"address"`
// 	Email        string `json:"email"`
// 	ContactNo    string `json:"contactNo"`
// 	Status       string `json:"status"`
// 	LastModified string `json:"lastModifed"`
// }
