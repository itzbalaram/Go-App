package interfaces

import "ecomm/web-service-gin/models"

type ICustomer interface {
	GetAll() ([]models.Customer, error)
	IfExists(email string) error
	Get(string) (*models.Product, error)
	Create(*models.Product) (interface{}, error)
	Delete(string) (interface{}, error)
}
