package database

import (
	// "ecomm/web-service-gin/interfaces"
	"ecomm/web-service-gin/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	ERROR_ORDER_EXISTS = errors.New("Order already exists with the given product id")
)

type OrderDB struct {
	Client interface{}
}

func (o *OrderDB) GetAll() ([]models.Order, error) {
	orders := []models.Order{}
	// Get all records
	result := o.Client.(*gorm.DB).Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (o *OrderDB) Get(id string) (*models.Order, error) {
	order := &models.Order{}
	result := o.Client.(*gorm.DB).First(order, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (o *OrderDB) Create(order *models.Order) (interface{}, error) {
	o.Client.(*gorm.DB).AutoMigrate(&models.Order{})
	result := o.Client.(*gorm.DB).Create(order)
	if result.Error != nil {
		fmt.Println("------------->", result.Error)
		return nil, result.Error
	}
	return order, nil
}

func (o *OrderDB) Delete(id string) (interface{}, error) {
	result := o.Client.(*gorm.DB).Delete(&models.Order{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.RowsAffected, nil
}

// func (o *OrderDB) IfExists(OrderId string) error {

// 	filter := map[string]interface{}{}
// 	filter["orderId"] = OrderId
// 	result := o.Client.(*gorm.DB).Model(&models.Order{}).First(&filter)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	if result.RowsAffected > 0 {
// 		return ERROR_ORDER_EXISTS
// 	}
// 	return nil
// }
