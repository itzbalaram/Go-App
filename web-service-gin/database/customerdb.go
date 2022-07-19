package database

import (
	// "ecomm/web-service-gin/interfaces"
	"ecomm/web-service-gin/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	ERROR_CUSTOMER_EXISTS = errors.New("Customer already exists with the given product id")
)

type CustomerDB struct {
	Client interface{}
}

func (c *CustomerDB) GetAll() ([]models.Customer, error) {
	customers := []models.Customer{}
	// Get all records
	result := c.Client.(*gorm.DB).Find(&customers)

	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (c *CustomerDB) Get(id string) (*models.Customer, error) {
	customer := &models.Customer{}
	result := c.Client.(*gorm.DB).First(customer, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return customer, nil
}

func (c *CustomerDB) Create(customer *models.Customer) (interface{}, error) {
	c.Client.(*gorm.DB).AutoMigrate(&models.Customer{})
	result := c.Client.(*gorm.DB).Create(customer)
	if result.Error != nil {
		fmt.Println("------------->", result.Error)
		return nil, result.Error
	}
	return customer, nil
}

func (c *CustomerDB) Delete(id string) (interface{}, error) {
	result := c.Client.(*gorm.DB).Delete(&models.Customer{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.RowsAffected, nil
}

func (c *CustomerDB) IfExists(Mobile string) error {

	filter := map[string]interface{}{}
	filter["mobile"] = Mobile
	result := c.Client.(*gorm.DB).Model(&models.Customer{}).First(&filter)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected > 0 {
		return ERROR_CUSTOMER_EXISTS
	}
	return nil
}
