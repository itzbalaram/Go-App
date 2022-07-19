package interfaces

import "ecomm/web-service-gin/models"

type IProduct interface {
	GetAll() ([]models.Product, error)
	IfExists(email string) error
	Get(string) (*models.Product, error)
	Create(*models.Product) (interface{}, error)
	Delete(string) (interface{}, error)
}
