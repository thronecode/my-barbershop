package service

import (
	"backend/utils"
)

// IService is the interface that defines the methods that the repository should implement
type IService interface {
	List(params *utils.RequestParams) (*Pag, error)
	Get(id *int) (*Service, error)
	GetByName(name *string) (*Service, error)
	Add(ser *Service) (*int, error)
	Update(id *int, ser *Service) error
	Delete(id *int) error
	AddPriceHistory(id *int, price *float64) error
}
