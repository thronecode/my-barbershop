package postgres

import (
	"github.com/thronecode/my-barbershop/backend/internal/config/database"
	"github.com/thronecode/my-barbershop/backend/internal/infra/barber"
	"github.com/thronecode/my-barbershop/backend/internal/sorry"
	"github.com/thronecode/my-barbershop/backend/internal/utils"

	"database/sql"
	"errors"
	"time"
)

type PGBarber struct {
	DB *database.DBTransaction
}

// List returns a paginated list of barbers
func (pg *PGBarber) List(params *utils.RequestParams) (*barber.Pag, error) {
	query := pg.DB.Builder.
		Select(utils.GetColumns(barber.Barber{}, &params.Total)...).
		From("t_barber bar").
		Where("deleted_at is null")

	var filters barber.FilterList
	err := params.ConvertFilters(&filters)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if filters.Name != nil {
		query = query.Where("lower(unaccent(name)) like lower(unaccent(?))", filters.Name)
	}

	barbers, next, count, err := utils.MakePaginatedList(barber.Barber{}, &query, params)
	if err != nil {
		return nil, sorry.Err(err)
	}

	return &barber.Pag{
		Data:  barbers.([]barber.Barber),
		Next:  next,
		Count: count,
	}, nil
}

// Get returns a barber by its id
func (pg *PGBarber) Get(id *int) (*barber.Barber, error) {
	var bar barber.Barber

	err := pg.DB.Builder.
		Select(utils.GetColumns(barber.Barber{}, nil)...).
		From("t_barber bar").
		Where("id = ?", id).
		Where("deleted_at is null").
		Scan(&bar.ID, &bar.Name, &bar.PhotoURL, &bar.CommissionRate, &bar.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, sorry.Err(err)
	}

	return &bar, nil
}

// Add adds a new barber
func (pg *PGBarber) Add(bar *barber.Barber) (*int, error) {
	var id int

	err := pg.DB.Builder.
		Insert("t_barber").
		Columns("name", "photo_url", "commission_rate").
		Values(bar.Name, bar.PhotoURL, bar.CommissionRate).
		Suffix("RETURNING id").
		Scan(&id)
	if err != nil {
		return nil, sorry.Err(err)
	}

	return &id, nil
}

// Update updates a barber's information
func (pg *PGBarber) Update(id *int, bar *barber.Barber) error {
	_, err := pg.DB.Builder.
		Update("t_barber").
		Set("name", bar.Name).
		Set("photo_url", bar.PhotoURL).
		Set("commission_rate", bar.CommissionRate).
		Where("id = ?", id).
		Exec()
	if err != nil {
		return sorry.Err(err)
	}

	return nil
}

// Delete deletes a barber
func (pg *PGBarber) Delete(id *int) error {
	_, err := pg.DB.Builder.
		Update("t_barber").
		Set("deleted_at", time.Now()).
		Where("id = ?", id).
		Where("deleted_at is null").
		Exec()
	if err != nil {
		return sorry.Err(err)
	}

	return nil
}

// AddCheckin adds a new check-in for a barber
func (pg *PGBarber) AddCheckin(checkin *barber.Checkin) (*int, error) {
	var id int

	err := pg.DB.Builder.
		Insert("t_barber_checkin").
		Columns("barber_id", "date_time").
		Values(checkin.BarberID, checkin.DateTime).
		Suffix("RETURNING id").
		Scan(&id)
	if err != nil {
		return nil, sorry.Err(err)
	}

	return &id, nil
}

// GetCheckins returns a list of check-ins for a barber
func (pg *PGBarber) GetCheckins(barberID *int, params *utils.RequestParams) (*barber.PagCheckin, error) {
	query := pg.DB.Builder.
		Select(utils.GetColumns(&barber.Checkin{}, nil)...).
		From("t_barber_checkin bch").
		Where("barber_id = ?", barberID)

	var filters barber.FilterCheckinList
	err := params.ConvertFilters(&filters)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if filters.InitialDate != nil {
		query = query.Where("date_time >= ?", filters.InitialDate)
	}

	if filters.FinalDate != nil {
		query = query.Where("date_time <= ?", filters.FinalDate)
	}

	checkins, next, count, err := utils.MakePaginatedList(barber.Checkin{}, &query, params)
	if err != nil {
		return nil, sorry.Err(err)
	}

	return &barber.PagCheckin{
		Data:  checkins.([]barber.Checkin),
		Next:  next,
		Count: count,
	}, nil
}

// AddService adds a new service for a barber
func (pg *PGBarber) AddService(barberID *int, services []int) error {
	insert := pg.DB.Builder.
		Insert("t_barber_service").
		Columns("barber_id", "service_id")

	for i := range services {
		insert = insert.Values(barberID, services[i])
	}

	_, err := insert.Exec()
	if err != nil {
		return sorry.Err(err)
	}

	return nil
}

// DeleteService deletes a service from a barber
func (pg *PGBarber) DeleteService(barberID, serviceID *int) error {
	_, err := pg.DB.Builder.
		Update("t_barber_service").
		Set("deleted_at", time.Now()).
		Where("service_id = ?", serviceID).
		Where("barber_id = ?", barberID).
		Where("deleted_at is null").
		Exec()
	if err != nil {
		return sorry.Err(err)
	}

	return nil
}
