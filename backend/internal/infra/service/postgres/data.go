package postgres

import (
	"github.com/thronecode/my-barbershop/backend/internal/config/database"
	"github.com/thronecode/my-barbershop/backend/internal/infra/service"
	"github.com/thronecode/my-barbershop/backend/internal/sorry"
	"github.com/thronecode/my-barbershop/backend/internal/utils"

	"database/sql"
	"errors"
	"time"
)

type PGService struct {
	DB *database.DBTransaction
}

// List returns a paginated list of services
func (pg *PGService) List(params *utils.RequestParams) (*service.Pag, error) {
	query := pg.DB.Builder.
		Select(utils.GetColumns(service.Service{}, &params.Total)...).
		From("t_service ser").
		Where("deleted_at is null")

	var filters service.FilterList
	err := params.ConvertFilters(&filters)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if filters.Name != nil {
		query = query.Where("lower(unaccent(name)) like lower(unaccent(?))", filters.Name)
	}
	if len(filters.Kinds) > 0 {
		query = query.Where("kinds @> ?", filters.Kinds)
	}
	if filters.IsCombo != nil {
		query = query.Where("is_combo = ?", filters.IsCombo)
	}
	if filters.BarberID != nil {
		query = query.Join("t_barber_service bser on ser.id = bser.service_id and bser.barber_id = ?", filters.BarberID)
	}

	services, next, count, err := utils.MakePaginatedList(service.Service{}, &query, params)
	if err != nil {
		return nil, sorry.Err(err)
	}

	return &service.Pag{
		Data:  services.([]service.Service),
		Next:  next,
		Count: count,
	}, nil
}

// Get returns a service by its id
func (pg *PGService) Get(id *int) (*service.Service, error) {
	var ser service.Service

	err := pg.DB.Builder.
		Select(utils.GetColumns(&service.Service{}, nil)...).
		From("t_service ser").
		Where("id = ?", id).
		Where("deleted_at is null").
		Scan(&ser.ID, &ser.Name, &ser.Description, &ser.Duration, &ser.Price, &ser.CommissionRate, &ser.IsCombo, &ser.Kinds, &ser.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, sorry.Err(err)
	}

	return &ser, nil
}

// GetByName returns a service by its name
func (pg *PGService) GetByName(name *string) (*service.Service, error) {
	var ser service.Service

	err := pg.DB.Builder.
		Select(utils.GetColumns(service.Service{}, nil)...).
		From("t_service ser").
		Where("name = ?", name).
		Where("deleted_at is null").
		Scan(&ser.ID, &ser.Name, &ser.Description, &ser.Duration, &ser.Price, &ser.CommissionRate, &ser.IsCombo, &ser.Kinds, &ser.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, sorry.Err(err)
	}

	return &ser, nil
}

// Add adds a new service
func (pg *PGService) Add(ser *service.Service) (*int, error) {
	var id int

	err := pg.DB.Builder.
		Insert("t_service").
		Columns("name", "description", "duration", "price", "commission_rate", "is_combo", "kinds").
		Values(ser.Name, ser.Description, &ser.Duration, ser.Price, ser.CommissionRate, ser.IsCombo, ser.Kinds).
		Suffix("RETURNING id").
		Scan(&id)
	if err != nil {
		return nil, sorry.Err(err)
	}

	return &id, nil
}

// Update updates a service
func (pg *PGService) Update(id *int, ser *service.Service) error {
	_, err := pg.DB.Builder.
		Update("t_service").
		Set("name", ser.Name).
		Set("description", ser.Description).
		Set("duration", ser.Duration).
		Set("price", ser.Price).
		Set("commission_rate", ser.CommissionRate).
		Set("is_combo", ser.IsCombo).
		Set("kinds", ser.Kinds).
		Where("id = ?", id).
		Exec()
	if err != nil {
		return sorry.Err(err)
	}

	return nil
}

// Delete deletes a service
func (pg *PGService) Delete(id *int) error {
	_, err := pg.DB.Builder.
		Update("t_service").
		Set("deleted_at", time.Now()).
		Where("id = ?", id).
		Where("deleted_at is null").
		Exec()
	if err != nil {
		return sorry.Err(err)
	}

	return nil
}

// AddPriceHistory adds a new price history for a service
func (pg *PGService) AddPriceHistory(serviceID *int, price *float64) error {
	_, err := pg.DB.Builder.
		Insert("t_service_price_history").
		Columns("service_id", "price", "date_time").
		Values(serviceID, price, time.Now()).
		Exec()
	if err != nil {
		return sorry.Err(err)
	}

	return nil
}
