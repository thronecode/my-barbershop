package admin

import (
	"backend/config/database"
	"backend/infra/admin"
	"backend/infra/admin/postgres"
	"backend/utils"
)

type repository struct {
	pg *postgres.PGAdmin
}

// New creates a new instance of the repository
func New(tx *database.DBTransaction) IAdmin {
	return &repository{pg: &postgres.PGAdmin{DB: tx}}
}

// List returns a paginated list of admins
func (r *repository) List(params *utils.RequestParams) (*Pag, error) {
	data, err := r.pg.List(params)
	if err != nil {
		return nil, err
	}

	res := &Pag{
		Data:  make([]Admin, len(data.Data)),
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

// Get returns an admin by its id
func (r *repository) Get(id *int) (*Admin, error) {
	data, err := r.pg.Get(id)
	if err != nil {
		return nil, err
	}

	res := new(Admin)
	if err = utils.ConvertStruct(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetByUsername returns an admin by its username
func (r *repository) GetByUsername(username *string) (*Admin, error) {
	data, err := r.pg.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	res := new(Admin)
	if err = utils.ConvertStruct(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Add adds a new admin
func (r *repository) Add(adm *Admin) (*int, error) {
	data := new(admin.Admin)
	if err := utils.ConvertStruct(adm, data); err != nil {
		return nil, err
	}

	return r.pg.Add(data)
}

// Update updates the password of an admin
func (r *repository) Update(id *int, password *string) error {
	return r.pg.Update(id, password)
}

// Delete deletes an admin
func (r *repository) Delete(id *int) error {
	return r.pg.Delete(id)
}
