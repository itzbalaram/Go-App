package interfaces

import "ecomm/web-service-gin/models"

type ICustomer interface {
	GetAll() ([]models.Customer, error)
	Get(string) (*models.Customer, error)
	Create(*models.Customer) (interface{}, error)
	Delete(string) (interface{}, error)
	IfExists(mobile string) error
}
