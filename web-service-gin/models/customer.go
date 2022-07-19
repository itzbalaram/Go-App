package models

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

// product represents data about a record product.
type Customer struct {
	gorm.Model // adds ID, created_at etc.

	// CustomerId   uint   `json:"customerId" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"index"`
	Mobile  string `json:"mobile"`
	Address string `json:"address"`
	// LastModified string `json:"lastModifed"`
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("invalid data for the field:Name")
	}
	if c.Mobile == "" {
		return fmt.Errorf("invalid data for the field:Mobile")
	}
	return nil
}

func (c *Customer) ToJsonByte() ([]byte, error) {
	return json.Marshal(c)
}

// products slice to seed record product data.
// var Customers = []Customer{
// 	{Name: "Redmi K20 Pro", Category: "Smartphone", Status: "Active", Price: 29999.99},
// 	{Name: "IPhone 13 pro", Category: "Smartphone", Status: "Active", Price: 75000.99},
// 	{Name: "HP-Pavilion 14", Category: "Laptop", Status: "Active", Price: 65000.99},
// 	{Name: "HP-Probook 16", Category: "Laptop", Status: "Active", Price: 85000.99},
// }
