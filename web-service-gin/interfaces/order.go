package interfaces

import "ecomm/web-service-gin/models"

type IOrder interface {
	GetAll() ([]models.Order, error)
	Get(string) (*models.Order, error)
	Create(*models.Order) (interface{}, error)
	Delete(string) (interface{}, error)
	// IfExists(id string) error
	// ProductDetails(string) (*models.Product, error)
}
