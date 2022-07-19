package database

import (
	// "ecomm/web-service-gin/interfaces"
	"ecomm/web-service-gin/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	ERROR_PRODUCT_EXISTS = errors.New("Product already exists with the given product id")
)

type ProductDB struct {
	Client interface{}
}

func (c *ProductDB) GetAll() ([]models.Product, error) {
	products := []models.Product{}
	// Get all records
	result := c.Client.(*gorm.DB).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (c *ProductDB) Get(id string) (*models.Product, error) {
	product := &models.Product{}
	result := c.Client.(*gorm.DB).First(product, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (c *ProductDB) Create(product *models.Product) (interface{}, error) {
	c.Client.(*gorm.DB).AutoMigrate(&models.Product{})
	result := c.Client.(*gorm.DB).Create(product)
	if result.Error != nil {
		fmt.Println("------------->", result.Error)
		return nil, result.Error
	}
	return product, nil
}

func (c *ProductDB) Delete(id string) (interface{}, error) {
	result := c.Client.(*gorm.DB).Delete(&models.Product{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.RowsAffected, nil
}

func (c *ProductDB) IfExists(Name string) error {

	filter := map[string]interface{}{}
	filter["name"] = Name
	result := c.Client.(*gorm.DB).Model(&models.Product{}).First(&filter)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected > 0 {
		return ERROR_PRODUCT_EXISTS
	}
	return nil
}
