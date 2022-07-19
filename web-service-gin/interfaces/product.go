package interfaces

import "ecomm/web-service-gin/models"

type IProduct interface {
	GetAll() ([]models.Product, error)
	Get(string) (*models.Product, error)
	Create(*models.Product) (interface{}, error)
	Delete(string) (interface{}, error)
	IfExists(name string) error
}
