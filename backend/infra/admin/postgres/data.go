package postgres

import (
	"backend/config/database"
	"backend/infra/admin"
	"backend/sorry"
	"backend/utils"
	"database/sql"

	"errors"
)

type PGAdmin struct {
	DB *database.DBTransaction
}

// List returns a paginated list of admins
func (pg *PGAdmin) List(params *utils.RequestParams) (*admin.PagAdmin, error) {
	query := pg.DB.Builder.
		Select(utils.GetColumns(admin.Admin{}, &params.Total)...).
		From("t_admin adm")

	var filters admin.FilterListAdmins
	err := params.ConvertFilters(&filters)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if filters.Username != nil {
		query = query.Where("lower(unaccent(username)) like lower(unaccent(?))", filters.Username)
	}

	admins, next, count, err := utils.MakePaginatedList(admin.Admin{}, &query, params)
	if err != nil {
		return nil, sorry.Err(err)
	}

	return &admin.PagAdmin{
		Data:  admins.([]admin.Admin),
		Next:  next,
		Count: count,
	}, nil
}

// Get returns an admin by its id
func (pg *PGAdmin) Get(id *int) (*admin.Admin, error) {
	var adm admin.Admin

	err := pg.DB.Builder.
		Select("adm.id as id", "adm.username as username").
		From("t_admin adm").
		Where("id = ?", id).
		Scan(&adm.ID, &adm.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, sorry.Err(err)
	}

	return &adm, nil
}

// GetByUsername returns an admin by its username
func (pg *PGAdmin) GetByUsername(username *string) (*admin.Admin, error) {
	var adm admin.Admin

	err := pg.DB.Builder.
		Select("adm.id as id", "adm.username as username", "adm.password as password").
		From("t_admin adm").
		Where("username = ?", username).
		Scan(&adm.ID, &adm.Username, &adm.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, sorry.Err(err)
	}

	return &adm, nil
}

// Add adds a new admin
func (pg *PGAdmin) Add(adm *admin.Admin) (*int, error) {
	var id int

	err := pg.DB.Builder.
		Insert("t_admin").
		Columns("username", "password").
		Values(adm.Username, adm.Password).
		Suffix("RETURNING id").
		Scan(&id)
	if err != nil {
		return nil, sorry.Err(err)
	}

	return &id, nil
}

// Update updates an admin
func (pg *PGAdmin) Update(id *int, password *string) error {
	_, err := pg.DB.Builder.
		Update("t_admin").
		Set("password", password).
		Where("id = ?", id).
		Exec()
	if err != nil {
		return sorry.Err(err)
	}

	return nil
}

// Delete deletes an admin
func (pg *PGAdmin) Delete(id *int) error {
	_, err := pg.DB.Builder.
		Delete("t_admin").
		Where("id = ?", id).
		Exec()
	if err != nil {
		return sorry.Err(err)
	}

	return nil
}
