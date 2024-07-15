package barber

import (
	"github.com/thronecode/my-barbershop/backend/internal/config/database"
	"github.com/thronecode/my-barbershop/backend/internal/infra/barber"
	"github.com/thronecode/my-barbershop/backend/internal/infra/barber/postgres"
	"github.com/thronecode/my-barbershop/backend/internal/utils"
)

type repository struct {
	pg *postgres.PGBarber
}

// New creates a new instance of the repository
func New(tx *database.DBTransaction) IBarber {
	return &repository{pg: &postgres.PGBarber{DB: tx}}
}

// List returns a paginated list of barbers
func (r *repository) List(params *utils.RequestParams) (*Pag, error) {
	data, err := r.pg.List(params)
	if err != nil {
		return nil, err
	}

	res := &Pag{
		Data:  make([]Barber, len(data.Data)),
		Next:  data.Next,
		Count: data.Count,
	}

	for i := range data.Data {
		if err = utils.ConvertStruct(&data.Data[i], &res.Data[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// Get returns a barber by its id
func (r *repository) Get(id *int) (*Barber, error) {
	data, err := r.pg.Get(id)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	res := new(Barber)
	if err = utils.ConvertStruct(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Add adds a new barber
func (r *repository) Add(bar *Barber) (*int, error) {
	data := new(barber.Barber)
	if err := utils.ConvertStruct(bar, data); err != nil {
		return nil, err
	}

	return r.pg.Add(data)
}

// Update updates a barber's information
func (r *repository) Update(id *int, bar *Barber) error {
	data := new(barber.Barber)
	if err := utils.ConvertStruct(bar, data); err != nil {
		return err
	}

	return r.pg.Update(id, data)
}

// Delete deletes a barber
func (r *repository) Delete(id *int) error {
	return r.pg.Delete(id)
}

// AddCheckin adds a new check-in for a barber
func (r *repository) AddCheckin(checkin *Checkin) (*int, error) {
	data := new(barber.Checkin)
	if err := utils.ConvertStruct(checkin, data); err != nil {
		return nil, err
	}

	return r.pg.AddCheckin(data)
}

// GetCheckins returns a list of check-ins for a barber
func (r *repository) GetCheckins(barberID *int, params *utils.RequestParams) (*PagCheckin, error) {
	data, err := r.pg.GetCheckins(barberID, params)
	if err != nil {
		return nil, err
	}

	res := &PagCheckin{
		Data:  make([]Checkin, len(data.Data)),
		Next:  data.Next,
		Count: data.Count,
	}

	for i := range data.Data {
		if err = utils.ConvertStruct(&data.Data[i], &res.Data[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// AddService adds a new service for a barber
func (r *repository) AddService(barberID *int, services []int) error {
	return r.pg.AddService(barberID, services)
}

// DeleteService deletes a service from a barber
func (r *repository) DeleteService(barberID, serviceID *int) error {
	return r.pg.DeleteService(barberID, serviceID)
}
