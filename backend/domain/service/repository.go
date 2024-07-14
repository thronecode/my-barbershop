package service

import (
	"backend/config/database"
	"backend/infra/service"
	"backend/infra/service/postgres"
	"backend/utils"
)

type repository struct {
	pg *postgres.PGService
}

// New creates a new instance of the repository
func New(tx *database.DBTransaction) IService {
	return &repository{pg: &postgres.PGService{DB: tx}}
}

// List returns a paginated list of services
func (r *repository) List(params *utils.RequestParams) (*Pag, error) {
	data, err := r.pg.List(params)
	if err != nil {
		return nil, err
	}

	res := &Pag{
		Data:  make([]Service, len(data.Data)),
		Next:  data.Next,
		Count: data.Count,
	}

	for i := range data.Data {
		if err := utils.ConvertStruct(&data.Data[i], &res.Data[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// Get returns a service by its id
func (r *repository) Get(id *int) (*Service, error) {
	ser, err := r.pg.Get(id)
	if err != nil {
		return nil, err
	}

	if ser == nil {
		return nil, nil
	}

	res := new(Service)
	if err := utils.ConvertStruct(ser, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetByName returns a service by its name
func (r *repository) GetByName(name *string) (*Service, error) {
	ser, err := r.pg.GetByName(name)
	if err != nil {
		return nil, err
	}

	if ser == nil {
		return nil, nil
	}

	res := new(Service)
	if err := utils.ConvertStruct(ser, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Add adds a new service
func (r *repository) Add(ser *Service) (*int, error) {
	data := new(service.Service)
	if err := utils.ConvertStruct(ser, data); err != nil {
		return nil, err
	}

	return r.pg.Add(data)
}

// Update updates a service
func (r *repository) Update(id *int, ser *Service) error {
	data := new(service.Service)
	if err := utils.ConvertStruct(ser, data); err != nil {
		return err
	}

	return r.pg.Update(id, data)
}

// Delete deletes a service
func (r *repository) Delete(id *int) error {
	return r.pg.Delete(id)
}

// AddPriceHistory adds a new price history
func (r *repository) AddPriceHistory(serviceID *int, price *float64) error {
	return r.pg.AddPriceHistory(serviceID, price)
}
