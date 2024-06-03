package barber

import (
	"backend/config/database"
	"backend/infra/barber"
	"backend/infra/barber/postgres"
	"backend/utils"
)

type repository struct {
	pg *postgres.PGBarber
}

// New creates a new instance of the repository
func New(tx *database.DBTransaction) IBarber {
	return &repository{pg: &postgres.PGBarber{DB: tx}}
}

// List returns a paginated list of barbers
func (r *repository) List(params *utils.RequestParams) (*barber.Pag, error) {
	return r.pg.List(params)
}

// Get returns a barber by its id
func (r *repository) Get(id *int) (*barber.Barber, error) {
	return r.pg.Get(id)
}

// Add adds a new barber
func (r *repository) Add(bar *barber.Barber) (*int, error) {
	return r.pg.Add(bar)
}

// Update updates a barber's information
func (r *repository) Update(id *int, bar *barber.Barber) error {
	return r.pg.Update(id, bar)
}

// Delete deletes a barber
func (r *repository) Delete(id *int) error {
	return r.pg.Delete(id)
}

// AddCheckin adds a new check-in for a barber
func (r *repository) AddCheckin(checkin *barber.Checkin) (*int, error) {
	return r.pg.AddCheckin(checkin)
}

// GetCheckins returns a list of check-ins for a barber
func (r *repository) GetCheckins(barberID *int, params *utils.RequestParams) (*barber.PagCheckin, error) {
	return r.pg.GetCheckins(barberID, params)
}
